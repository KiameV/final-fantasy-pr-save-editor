package misc

import (
	"github.com/aarzilli/nucular"
	"math"
	"pr_save_editor/global"
	"pr_save_editor/models"
	"pr_save_editor/ui"
	"pr_save_editor/ui/widgets"
)

type miscUI struct {
	yLast int
}

func NewUI() ui.UI {
	return &miscUI{}
}

func (u *miscUI) Draw(w *nucular.Window) {
	m := models.GetMisc()

	w.Row(24).Static(200, 10, 200, 10, 250)
	w.PropertyInt("GP:", 0, &m.GP, 16777216, 1000, 0)
	w.Spacing(1)
	w.PropertyInt("Steps:", 0, &m.Steps, 16777216, 1000, 0)
	w.Spacing(1)
	if global.Six == global.GetSaveType() {
		w.PropertyInt("Cursed Shield Fight Count:", 0, &m.CursedShieldFightCount, 255, 1, 0)
	} else {
		w.Spacing(1)
	}

	w.Row(5).Static(0)

	w.Row(24).Static(200, 10, 200, 10, 250)
	w.PropertyInt("Save Count:", 0, &m.NumberOfSaves, 16777216, 1, 0)
	w.Spacing(1)
	w.PropertyInt("Battle Count:", 0, &m.BattleCount, 16777216, 1, 0)
	w.Spacing(1)
	w.PropertyInt("Escape Count:", 0, &m.EscapeCount, 16777216, 1, 0)

	w.Row(5).Static(0)

	w.Row(24).Static(300)
	w.CheckboxText("Is Complete Flag", &m.IsCompleteFlag)

	w.Row(24).Static(250, 400)
	_ = w.PropertyInt("Chests Opened", 0, &m.OpenedChestCount, 214, 1, 0)

	if global.GetSaveType() == global.One {
		w.Row(5).Static()
		w.Row(24).Static(500)
		w.Label("Need help figuring out which is Fire, Ice, Wind, and Earth", "LC")

		w.Row(24).Static(100)
		w.CheckboxText("Crystal 1", &m.OwnedCrystals[0])
		w.Row(24).Static(100)
		w.CheckboxText("Crystal 2", &m.OwnedCrystals[1])
		w.Row(24).Static(100)
		w.CheckboxText("Crystal 3", &m.OwnedCrystals[2])
		w.Row(24).Static(100)
		w.CheckboxText("Crystal 4", &m.OwnedCrystals[3])
	}

	w.Row(5).Static()

	w.Row(24).Static(200)
	w.Label("Played Time:", "LC")
	hours, minutes := widgets.GetTime(int(m.PlayTime))
	w.Row(24).Static(200, 200)
	b1 := w.PropertyInt("Hours", 0, &hours, math.MaxInt, 1, 0)
	b2 := w.PropertyInt("Minutes", 0, &minutes, 59, 1, 0)
	if b1 || b2 {
		m.PlayTime = float64(hours*3600 + minutes*60)
	}
}

func (u *miscUI) Refresh() {

}

func (u *miscUI) Name() string {
	return "Misc"
}

func (u *miscUI) Behavior() ui.Behavior {
	return ui.Show
}
