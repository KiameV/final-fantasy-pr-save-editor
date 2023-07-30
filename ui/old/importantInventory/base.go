package importantInventory

import (
	"fmt"
	"strings"

	"github.com/aarzilli/nucular"
	"pr_save_editor/consts"
	"pr_save_editor/consts/ffv"
	"pr_save_editor/consts/ffvi"
	"pr_save_editor/global"
	"pr_save_editor/models"
	"pr_save_editor/ui/old"
)

type importantInventoryUI struct {
	ids     []*nucular.TextEditor
	counts  []*nucular.TextEditor
	helper  nucular.TextEditor
	lookups map[int]string
}

func NewUI() old.UI {
	inv := &importantInventoryUI{
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

func (u *importantInventoryUI) Draw(w *nucular.Window) {
	inv := models.GetImportantInventory()

	// Items
	w.Row(400).Static(230, 230)

	if sw := w.GroupBegin("iitems", 0); sw != nil {
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
			if u.lookups != nil {
				if s, found := u.lookups[r.ItemID]; found {
					sw.Row(20).Static(210)
					sw.Label(s, "LT")
					sw.Row(4).Static(1)
					sw.Spacing(1)
				}
			}
		}
		sw.GroupEnd()
	}

	switch global.GetSaveType() {
	case global.Five, global.Six:
		if sw := w.GroupBegin("iitextboxes", 0); sw != nil {
			sw.Row(20).Static(200)
			sw.Label("Important Items", "LC")
			sw.Row(300).Static(200)
			u.helper.Edit(sw)
			sw.GroupEnd()
		}
	default:
		w.Spacing(2)
	}
}

func (u *importantInventoryUI) Refresh() {
	switch global.GetSaveType() {
	case global.Five:
		u.lookups = ffv.ImportantLookup
		initReadOnlyText(&u.helper, ffv.ImportantItems)
	case global.Six:
		u.lookups = ffvi.ImportantLookup
		initReadOnlyText(&u.helper, ffvi.ImportantItems)
	default:
		u.lookups = nil
	}
}

func (u *importantInventoryUI) Name() string {
	return "Important Inventory"
}

func (u *importantInventoryUI) Behavior() old.Behavior {
	return old.Show
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