package importantInventory

import (
	"github.com/aarzilli/nucular"
	"pr_save_editor/models"
	"pr_save_editor/ui"
)

type importantInventoryUI struct {
	ids    []*nucular.TextEditor
	counts []*nucular.TextEditor
}

func NewUI() ui.UI {
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

	// Finder
	// + widgets.DrawItemFinder(w, 540, 24)

	// Items
	w.Row(24).Static(100, 10, 100)
	w.Label("Item ID", "LC")
	w.Spacing(1)
	w.Label("Count", "LC")

	isFirstEmptyRow := true
	for _, r := range inv.GetRows() {
		if r.ItemID == 0 && r.Count == 0 {
			if isFirstEmptyRow {
				isFirstEmptyRow = false
			} else {
				continue
			}
		}

		w.Row(24).Static(100, 10, 100)
		w.PropertyInt("", 0, &r.ItemID, 999, 1, 0)
		w.Spacing(1)
		w.PropertyInt("", 0, &r.Count, 999, 1, 0)
	}
}

func (u *importantInventoryUI) Refresh() {

}

func (u *importantInventoryUI) Name() string {
	return "Important Inventory"
}

func (u *importantInventoryUI) Behavior() ui.Behavior {
	return ui.Show
}
