package inventory

import (
	"encoding/json"
	"github.com/aarzilli/nucular"
	"pr_save_editor/io"
	"pr_save_editor/models"
	"pr_save_editor/ui"
)

type inventoryUI struct {
	ids    []*nucular.TextEditor
	counts []*nucular.TextEditor
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
		inv   = models.GetInventory()
	)

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
	// widgets.DrawItemFinder(w, 540, 24)

	// Items
	w.Row(24).Static(100, 10, 100)
	w.Label("Item ID", "CC")
	w.Spacing(1)
	w.Label("Count", "CC")

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
		y += 24
		count += 2
	}
}

func (u *inventoryUI) Refresh() {

}

func (u *inventoryUI) Name() string {
	return "Inventory"
}

func (u *inventoryUI) Behavior() ui.Behavior {
	return ui.Show
}
