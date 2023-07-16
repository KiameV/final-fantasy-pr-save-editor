package inventory

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aarzilli/nucular"
	"pr_save_editor/consts"
	"pr_save_editor/consts/ffv"
	"pr_save_editor/consts/ffvi"
	"pr_save_editor/global"
	"pr_save_editor/io"
	"pr_save_editor/models"
	"pr_save_editor/ui"
	"pr_save_editor/ui/widgets"
)

type inventoryUI struct {
	ids     []*nucular.TextEditor
	counts  []*nucular.TextEditor
	lookups map[int]string

	itemsHelp   nucular.TextEditor
	weaponsHelp nucular.TextEditor
	armorHelp   nucular.TextEditor
	relicHelp   nucular.TextEditor
}

func NewUI() ui.UI {
	inv := &inventoryUI{
		ids:    make([]*nucular.TextEditor, 255),
		counts: make([]*nucular.TextEditor, 255),
	}
	for i := 0; i < 255; i++ {
		tb := &nucular.TextEditor{}
		inv.ids[i] = tb
		tb.Flags = nucular.EditField
		tb.SingleLine = true
		tb.Maxlen = 3

		tb = &nucular.TextEditor{}
		inv.counts[i] = tb
		tb.Flags = nucular.EditField
		tb.SingleLine = true
		tb.Maxlen = 3
	}
	return inv
}

func (u *inventoryUI) Draw(w *nucular.Window) {
	inv := models.GetInventory()

	// Top
	w.Row(24).Static(150, 10, 150, 10, 100, 10, 100)
	w.CheckboxText("Reset Sort Order", &inv.ResetSortOrder)
	w.Spacing(1)
	w.CheckboxText("Remove Duplicates", &inv.RemoveDuplicates)
	w.Spacing(1)
	if w.ButtonText("Save") {
		if result, err := json.Marshal(inv.Rows); err != nil {
			ui.DrawError = err
		} else if err = io.SaveInvFile(w, result); err != nil {
			ui.DrawError = err
		}
	}
	w.Spacing(1)
	if w.ButtonText("Load") {
		if data, err := io.OpenInvFileDialog(w); err != nil {
			ui.DrawError = err
		} else if err = json.Unmarshal(data, &inv.Rows); err != nil {
			ui.DrawError = err
		}
		inv.ResetSortOrder = true
	}

	w.Row(400).Static(230, 230, 230)

	if sw := w.GroupBegin("items", 0); sw != nil {
		// Items
		sw.Row(24).Static(100, 10, 100)
		sw.Label("Item ID", "CC")
		sw.Spacing(1)
		sw.Label("Count", "CC")

		isFirstEmptyRow := true
		for _, r := range inv.GetRows() {
			if r.ItemID == 0 && r.Count == 0 {
				if isFirstEmptyRow {
					isFirstEmptyRow = false
				} else {
					continue
				}
			}

			sw.Row(18).Static(100, 10, 100)
			sw.PropertyInt("", 0, &r.ItemID, 999, 1, 0)
			sw.Spacing(1)
			sw.PropertyInt("", 0, &r.Count, 999, 1, 0)
			if s, found := u.lookups[r.ItemID]; found {
				sw.Row(20).Static(210)
				sw.Label(s, "LT")
				sw.Row(4).Static(1)
				sw.Spacing(1)
			}
		}
		sw.GroupEnd()
	}

	switch global.GetSaveType() {
	case global.Five:
		if sw := w.GroupBegin("textboxes", 0); sw != nil {
			sw.Row(20).Static(200)
			sw.Label("Items", "LC")
			sw.Row(100).Static(200)
			u.itemsHelp.Edit(sw)

			sw.Row(20).Static(200)
			sw.Label("Weapons", "LC")
			sw.Row(100).Static(200)
			u.weaponsHelp.Edit(sw)

			sw.Row(20).Static(200)
			sw.Label("Armors", "LC")
			sw.Row(100).Static(200)
			u.armorHelp.Edit(sw)
			sw.GroupEnd()
		}
		widgets.AddItemFinder(w, ffv.Items, ffv.Weapons, ffv.Armors)
	case global.Six:
		if sw := w.GroupBegin("textboxes", 0); sw != nil {
			sw.Row(20).Static(200)
			sw.Label("Items", "LC")
			sw.Row(70).Static(200)
			u.itemsHelp.Edit(sw)

			sw.Row(20).Static(200)
			sw.Label("Weapons", "LC")
			sw.Row(70).Static(200)
			u.weaponsHelp.Edit(sw)

			sw.Row(20).Static(200)
			sw.Label("Armors", "LC")
			sw.Row(70).Static(200)
			u.armorHelp.Edit(sw)

			sw.Row(20).Static(200)
			sw.Label("Relics", "LC")
			sw.Row(70).Static(200)
			u.relicHelp.Edit(sw)
			sw.GroupEnd()
		}
		widgets.AddItemFinder(w, ffv.Items, ffv.Weapons, ffv.Armors)
	default:
		w.Spacing(2)
	}
}

func (u *inventoryUI) Refresh() {
	switch global.GetSaveType() {
	case global.Five:
		u.lookups = ffv.GeneralLookup
		initReadOnlyText(&u.itemsHelp, ffv.Items)
		initReadOnlyText(&u.weaponsHelp, ffv.Weapons)
		initReadOnlyText(&u.armorHelp, ffv.Armors)
	case global.Six:
		u.lookups = ffvi.GeneralLookup
		initReadOnlyText(&u.itemsHelp, ffvi.Items)
		initReadOnlyText(&u.weaponsHelp, ffvi.Weapons)
		initReadOnlyText(&u.armorHelp, ffvi.Armors)
		initReadOnlyText(&u.relicHelp, ffvi.Relics)
	default:
		u.lookups = nil
	}
}

func (u *inventoryUI) Name() string {
	return "Inventory"
}

func (u *inventoryUI) Behavior() ui.Behavior {
	return ui.Show
}

func initReadOnlyText(tb *nucular.TextEditor, items []*consts.Item) {
	tb.Flags = nucular.EditBox | nucular.EditMultiline
	tb.SingleLine = false
	tb.SelectAll()
	tb.DeleteSelection()
	sb := strings.Builder{}
	for _, i := range items {
		_, _ = sb.WriteString(fmt.Sprintf("%d - %s\n", i.ID, i.Name))
	}
	tb.Text([]rune(sb.String()))
	tb.Flags |= nucular.EditReadOnly
	tb.Flags |= nucular.EditNoHorizontalScroll
}
