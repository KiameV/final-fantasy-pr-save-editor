package inventory

import (
	"encoding/json"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/rect"
	"pr_save_editor/io"
	"pr_save_editor/models"
	"pr_save_editor/ui"
)

type inventoryUI struct {
	ids       []*nucular.TextEditor
	counts    []*nucular.TextEditor
	countLast int
	yLast     int
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
	var (
		y     int
		count = 4
	)

	// Top
	w.Row(u.yLast).SpaceBegin(u.countLast)
	w.LayoutSpacePush(rect.Rect{
		X: 0,
		Y: y,
		W: 150,
		H: 22,
	})
	w.CheckboxText("Reset Sort Order", &models.GetInventory().ResetSortOrder)

	w.LayoutSpacePush(rect.Rect{
		X: 160,
		Y: y,
		W: 150,
		H: 22,
	})
	w.CheckboxText("Remove Duplicates", &models.GetInventory().RemoveDuplicates)

	w.LayoutSpacePush(rect.Rect{
		X: 320,
		Y: y,
		W: 100,
		H: 22,
	})
	if w.ButtonText("Save") {
		if result, err := json.Marshal(models.GetInventory().Rows); err != nil {
			ui.DrawError = err
		} else if err = io.SaveInvFile(w, result); err != nil {
			ui.DrawError = err
		}
	}

	w.LayoutSpacePush(rect.Rect{
		X: 430,
		Y: y,
		W: 100,
		H: 22,
	})
	if w.ButtonText("Load") {
		if data, err := io.OpenInvFileDialog(w); err != nil {
			ui.DrawError = err
		} else if err = json.Unmarshal(data, &models.GetInventory().Rows); err != nil {
			ui.DrawError = err
		}
		models.GetInventory().ResetSortOrder = true
	}
	y += 24

	// Items
	w.LayoutSpacePush(rect.Rect{
		X: 0,
		Y: y,
		W: 75,
		H: 22,
	})
	w.Label("Item ID", "LC")

	w.LayoutSpacePush(rect.Rect{
		X: 85,
		Y: y,
		W: 75,
		H: 22,
	})
	w.Label("Count", "LC")
	y += 24

	for _, r := range models.GetInventoryRows() {
		if r.ItemID == 93 || r.ItemID == 198 || r.ItemID == 199 || r.ItemID == 200 {
			continue
		}
		w.LayoutSpacePush(rect.Rect{
			X: 0,
			Y: y,
			W: 75,
			H: 24,
		})
		w.PropertyInt("", 0, &r.ItemID, 999, 1, 0)

		w.LayoutSpacePush(rect.Rect{
			X: 85,
			Y: y,
			W: 75,
			H: 24,
		})
		w.PropertyInt("", 0, &r.Count, 999, 1, 0)
		y += 24
		count += 2
	}
	u.yLast = y

	// Finder
	u.countLast = count // + widgets.DrawItemFinder(w, 540, 24)
}

func (u *inventoryUI) Refresh() {

}

func (u *inventoryUI) Name() string {
	return "Inventory"
}

func (u *inventoryUI) Behavior() ui.Behavior {
	return ui.Show
}
