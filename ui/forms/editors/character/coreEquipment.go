package character

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/core/costs/ff6"
	"pixel-remastered-save-editor/models/finder"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Equipment struct {
		widget.BaseWidget
		game   global.Game
		inputs []*eqRow
	}
	eqRow struct {
		*inputs.IdEntry
		name string
	}
)

func NewCoreEquipment(game global.Game, c *core.Character) *Equipment {
	e := &Equipment{
		game:   game,
		inputs: make([]*eqRow, 0, len(c.Equipment.Values)+1),
	}
	e.ExtendBaseWidget(e)
	for i, j := range c.Equipment.Values {
		e.inputs = append(e.inputs, &eqRow{
			IdEntry: inputs.NewIdEntryWithDataWithHint(&j.ContentID, finder.Items),
		})
		switch i {
		case 0:
			e.inputs[i].name = "R. Hand:"
		case 1:
			e.inputs[i].name = "L. Hand:"
		case 2:
			e.inputs[i].name = "Body:"
		case 3:
			e.inputs[i].name = "Head:"
		default:
			e.inputs[i].name = "Arms:"
		}
	}
	if game.IsSix() {
		e.inputs = append(e.inputs, &eqRow{
			IdEntry: inputs.NewIdEntryWithDataWithHint(&c.Base.MagicStoneID, finder.New(ff6.Espers)),
			name:    "Esper:",
		})
	}
	return e
}

func (e *Equipment) CreateRenderer() fyne.WidgetRenderer {
	rows := container.NewVBox()
	itSearch := inputs.GetSearches().Items
	eqSearch := inputs.GetSearches().Equipment
	for _, i := range e.inputs {
		rows.Add(container.NewGridWithColumns(3, widget.NewLabel(i.name), i.Label, i.ID))
	}
	left := container.NewVScroll(rows)
	var right *fyne.Container
	if e.game.IsSix() {
		espSearch := inputs.NewSearch(ff6.Espers)
		right = container.NewGridWithColumns(5,
			eqSearch.Fields(), eqSearch.Filter(),
			espSearch.Fields(), espSearch.Filter(),
			itSearch.Fields())
	} else {
		right = container.NewGridWithColumns(4,
			eqSearch.Fields(), eqSearch.Filter(),
			itSearch.Fields(), itSearch.Filter())
	}
	return widget.NewSimpleRenderer(
		container.NewBorder(nil, nil, left, right, nil))
}
