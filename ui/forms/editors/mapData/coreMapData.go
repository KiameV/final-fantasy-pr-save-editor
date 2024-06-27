package mapData

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/finder"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	MapData struct {
		widget.BaseWidget
		player       []fyne.CanvasObject
		gps          []fyne.CanvasObject
		misc         []fyne.CanvasObject
		otherParties *fyne.Container
	}
)

func NewCore(data *core.MapData) *MapData {
	gpsMapID := inputs.NewIdEntryWithDataWithHint(&data.Gps.MapID, finder.Maps)
	miscMapID := inputs.NewIdEntryWithDataWithHint(&data.Map.MapID, finder.Maps)
	e := &MapData{
		player: []fyne.CanvasObject{
			inputs.NewLabeledEntry("Position X", inputs.NewFloatEntryWithData(&data.Player.Position.X), 3),
			inputs.NewLabeledEntry("Position Y", inputs.NewFloatEntryWithData(&data.Player.Position.Y), 3),
			inputs.NewLabeledEntry("Position Z", inputs.NewFloatEntryWithData(&data.Player.Position.Z), 3),
			inputs.NewLabeledEntry("Facing Direction", inputs.NewIntEntryWithData(&data.Player.Direction), 3),
		},
		gps: []fyne.CanvasObject{
			container.NewGridWithColumns(3, widget.NewLabel("Map ID"), gpsMapID.ID, gpsMapID.Label),
			inputs.NewLabeledEntry("Area ID", inputs.NewIntEntryWithData(&data.Gps.AreaID), 3),
			inputs.NewLabeledEntry("ID", inputs.NewIntEntryWithData(&data.Gps.GpsID), 3),
			inputs.NewLabeledEntry("Width", inputs.NewIntEntryWithData(&data.Gps.Width), 3),
			inputs.NewLabeledEntry("Height", inputs.NewIntEntryWithData(&data.Gps.Height), 3),
		},
		misc: []fyne.CanvasObject{
			container.NewGridWithColumns(3, widget.NewLabel("Map ID"), miscMapID.ID, miscMapID.Label),
			inputs.NewLabeledEntry("Point In", inputs.NewIntEntryWithData(&data.Map.PointIn), 3),
			inputs.NewLabeledEntry("Transportation ID", inputs.NewIntEntryWithData(&data.Map.TransportationID), 3),
			inputs.NewLabeledEntry("Carrying Hover Ship", widget.NewCheckWithData("", binding.BindBool(&data.Map.CarryingHoverShip)), 3),
		},
		otherParties: container.NewVBox(),
	}
	e.ExtendBaseWidget(e)
	if data.OtherParties != nil {
		if parties := *data.OtherParties; len(parties) > 0 && parties[0].PlayableCharacterCorpsId > -1 {
			for i, party := range parties {
				if p, loaded := e.loadParty(party); loaded {
					e.otherParties.Add(widget.NewLabel("Party " + strconv.Itoa(i+1)))
					mapID := inputs.NewIdEntryWithDataWithHint(&p.MapID, finder.Maps)
					c := container.NewVBox(
						container.NewGridWithColumns(3, widget.NewLabel("Map ID"), mapID.ID, mapID.Label),
						inputs.NewLabeledEntry("Point In", inputs.NewIntEntryWithData(&p.PointIn), 3))
					if p.PE != nil {
						c.Add(inputs.NewLabeledEntry("Position X", inputs.NewFloatEntryWithData(&p.PE.Position.X), 3))
						c.Add(inputs.NewLabeledEntry("Position Y", inputs.NewFloatEntryWithData(&p.PE.Position.Y), 3))
						c.Add(inputs.NewLabeledEntry("Position Z", inputs.NewFloatEntryWithData(&p.PE.Position.Z), 3))
					}
					e.otherParties.Add(container.NewPadded(c))
				}
			}
		}
	}
	return e
}

func (e *MapData) CreateRenderer() fyne.WidgetRenderer {
	search := inputs.GetSearches().Maps
	return widget.NewSimpleRenderer(container.NewBorder(nil, nil,
		container.NewGridWithColumns(2,
			container.NewVScroll(container.NewVBox(
				widget.NewLabel("Player"),
				container.NewPadded(container.NewVBox(e.player...)),
				widget.NewLabel("GPS"),
				container.NewPadded(container.NewVBox(e.gps...)),
				widget.NewLabel("Misc"),
				container.NewPadded(container.NewVBox(e.misc...)))),
			container.NewVScroll(e.otherParties),
		),
		container.NewGridWithColumns(2, search.Fields(), search.Filter())))
}

func (e *MapData) loadParty(party *save.OtherPartyData) (p *save.OtherPartyData, loaded bool) {
	pe, err := party.PlayerEntity()
	if err == nil {
		p = &save.OtherPartyData{
			MapID:                    party.MapID,
			PointIn:                  party.PointIn,
			PE:                       pe,
			PlayableCharacterCorpsId: party.PlayableCharacterCorpsId,
		}
		loaded = true
	}
	return
}
