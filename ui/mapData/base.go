package mapData

import (
	"math"

	"github.com/aarzilli/nucular"
	"pr_save_editor/global"
	"pr_save_editor/models"
	"pr_save_editor/ui"
)

type mapDataUI struct{}

func NewUI() ui.UI {
	return &mapDataUI{}
}

func (u *mapDataUI) Draw(w *nucular.Window) {
	m := models.GetMapData()

	w.Row(24).Static(200, 10, 200)
	w.PropertyInt("Map ID", 0, &m.MapID, math.MaxInt, 1, 0)
	w.Spacing(1)
	w.PropertyInt("PointIn", math.MinInt, &m.PointIn, math.MaxInt, 1, 0)

	w.Row(24).Static(200, 10, 200)
	w.PropertyInt("Transportation ID", 0, &m.TransportationID, math.MaxInt, 1, 0)
	w.Spacing(1)
	w.CheckboxText("Carrying Hover Ship", &m.CarryingHoverShip)

	if global.GetSaveType() == global.Six {
		w.Row(24).Static(410)
		w.PropertyInt("Playable Character Corps Id", 0, &m.PlayableCharacterCorpsID, math.MaxInt, 1000, 0)
	}

	w.Row(24).Static(300)
	w.Label("Player Position:", "LC")

	w.Row(24).Static(200, 10, 200, 10, 200)
	w.PropertyFloat("X", -math.MaxFloat32, &m.Player.X, math.MaxFloat32, 0.1, 0, 5)
	w.Spacing(1)
	w.PropertyFloat("Y", -math.MaxFloat32, &m.Player.Y, math.MaxFloat32, 0.1, 0, 5)
	w.Spacing(1)
	w.PropertyFloat("Z", -math.MaxFloat32, &m.Player.Z, math.MaxFloat32, 0.1, 0, 5)

	w.Row(24).Static(200)
	w.PropertyInt("Player Direction", 0, &m.PlayerDirection, math.MaxInt, 1, 0)

	w.Row(5).Static()

	w.Row(24).Static(300)
	w.Label("GPS Data:", "LC")

	w.Row(24).Static(200, 10, 200, 10, 250)
	w.PropertyInt("Map ID", 0, &m.Gps.MapID, math.MaxInt, 1, 0)
	w.Spacing(1)
	w.PropertyInt("Area ID", 0, &m.Gps.AreaID, math.MaxInt, 1, 0)
	w.Spacing(1)
	w.PropertyInt("GPS ID", 0, &m.Gps.GpsID, math.MaxInt, 1, 0)

	w.Row(24).Static(200, 10, 200)
	w.PropertyInt("Width", 0, &m.Gps.Width, math.MaxInt, 1, 0)
	w.Spacing(1)
	w.PropertyInt("Height", 0, &m.Gps.Height, math.MaxInt, 1, 0)
}

func (u *mapDataUI) Refresh() {

}

func (u *mapDataUI) Name() string {
	return "Map Data"
}

func (u *mapDataUI) Behavior() ui.Behavior {
	return ui.Show
}
