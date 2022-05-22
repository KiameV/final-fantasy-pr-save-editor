package transportation

import (
	"fmt"
	"github.com/aarzilli/nucular"
	"math"
	"pr_save_editor/global"
	"pr_save_editor/models"
	"pr_save_editor/ui"
)

type mapDataUI struct {
	yLast int
}

func NewUI() ui.UI {
	return &mapDataUI{}
}

func (u *mapDataUI) Draw(w *nucular.Window) {
	for i, t := range models.Transportations {
		if i > 0 {
			w.Row(4).Static()
		}

		w.Row(24).Static(50, 10, 100)
		w.Label(fmt.Sprintf("ID: %d", t.ID), "LC")
		w.Spacing(1)
		if w.CheckboxText("Enabled", &t.Enabled) {
			if t.Enabled {
				t.ForcedEnabled = true
				t.ForcedDisabled = false
			} else {
				t.ForcedEnabled = false
				t.ForcedDisabled = true
			}
		}

		if t.Enabled {
			w.Row(24).Static(200)
			w.PropertyInt("Map ID", math.MinInt, &t.MapID, math.MaxInt, 1, 0)

			w.Row(24).Static(200)
			w.PropertyFloat("Position X", -math.MaxFloat32, &t.Position.X, math.MaxFloat32, 0.1, 0, 5)
			w.Row(24).Static(200)
			w.PropertyFloat("Position Y", -math.MaxFloat32, &t.Position.Y, math.MaxFloat32, 0.1, 0, 5)
			w.Row(24).Static(200)
			w.PropertyFloat("Position Z", -math.MaxFloat32, &t.Position.Z, math.MaxFloat32, 0.1, 0, 5)

			w.Row(24).Static(200)
			w.PropertyInt("Direction", 0, &t.Direction, math.MaxInt, 1, 0)

			w.Row(24).Static(300, 10, 50)
			f := float64(t.TimeStampTicks)
			if w.PropertyFloat("Timestamp Ticks", 0, &f, math.MaxUint64, 1, 0, 100) {
				t.TimeStampTicks = uint64(f)
			}
			w.Spacing(1)
			if w.ButtonText("now") {
				t.TimeStampTicks = global.NowToTicks()
			}
		}
	}
}

func (u *mapDataUI) Refresh() {

}

func (u *mapDataUI) Name() string {
	return "Transportation"
}

func (u *mapDataUI) Behavior() ui.Behavior {
	return ui.Show
}
