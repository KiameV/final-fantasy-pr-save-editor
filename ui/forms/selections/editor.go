package selections

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/ui/forms/editors"
	"pixel-remastered-save-editor/ui/forms/editors/mapData"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Editor struct {
		widget.BaseWidget
		game     global.Game
		save     *core.Save
		renderer fyne.WidgetRenderer
	}
)

func NewEditor(game global.Game, save *core.Save) *Editor {
	s := &Editor{
		game: game,
		save: save,
	}
	s.ExtendBaseWidget(s)
	inputs.Load(game)

	tabs := container.NewAppTabs(
		container.NewTabItem("Characters", NewCharacters(s.save)),
		container.NewTabItem("Inventory", NewInventory(s.save)),
		container.NewTabItem("Party", editors.NewCoreParty(s.save.Party, s.save.Parties)),
		container.NewTabItem("Map Data", mapData.NewCore(s.save.Map)),
		container.NewTabItem("Transportation", editors.NewCoreTransportation(s.game, s.save.Transportations)),
		container.NewTabItem("Misc", NewMisc(s.save.Misc)),
	)
	// if s.save.Bestiary != nil {
	// 	tabs.Append(container.NewTabItem("Bestiary", NewBestiary(s.save.Bestiary)))
	// }
	if s.game.IsSix() {
		tabs.Append(container.NewTabItem("Espers", editors.NewEspers(s.save)))
	}
	s.renderer = widget.NewSimpleRenderer(tabs)
	return s
}

func (s *Editor) CreateRenderer() fyne.WidgetRenderer {
	return s.renderer
}
