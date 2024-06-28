package editors

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/finder"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Transportation struct {
		widget.BaseWidget
		tabs []*container.TabItem
	}
)

func NewCoreTransportation(game global.Game, transportations *core.Transportations) *Transportation {
	e := &Transportation{}
	e.ExtendBaseWidget(e)

	for _, t := range transportations.Transportations {
		if i, ok := e.createTabItem(game, t); ok {
			e.tabs = append(e.tabs, i)
		}
	}
	return e
}

func (e *Transportation) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewAppTabs(e.tabs...))
}

func (e *Transportation) createTabItem(game global.Game, transport *save.OwnedTransportation) (*container.TabItem, bool) {
	p := transport.Position
	if p == nil {
		return nil, false
	}
	name := strconv.Itoa(transport.ID)
	if game.IsSix() {
		if transport.ID == 4 {
			name = "Blackjack"
		}
		if transport.ID == 5 {
			name = "Falcon"
		}
	}
	in := inputs.NewIdEntryWithDataWithHint(&transport.MapID, finder.Maps)
	c := container.NewVBox(
		container.NewGridWithColumns(3, widget.NewLabel("Map ID"), in.ID, in.Label),
		inputs.NewLabeledEntry("Position X", inputs.NewFloatEntryWithData(&p.X), 3),
		inputs.NewLabeledEntry("Position Y", inputs.NewFloatEntryWithData(&p.Y), 3),
		inputs.NewLabeledEntry("Position Y", inputs.NewFloatEntryWithData(&p.Z), 3),
		inputs.NewLabeledEntry("Facing", inputs.NewIntEntryWithData(&transport.Direction), 3),
		inputs.NewLabeledEntry("Time Stamp Ticks", inputs.NewIntEntryWithData(&transport.TimeStampTicks), 3))
	return container.NewTabItem(name, container.NewBorder(nil, nil, c, nil)), true
}
