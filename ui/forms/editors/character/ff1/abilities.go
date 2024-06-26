package ff1

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/finder"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Abilities struct {
		widget.BaseWidget
		c         *core.Character
		abilities *fyne.Container
		body      *fyne.Container
		add       *widget.Button
	}
)

func NewAbilities(c *core.Character) *Abilities {
	e := &Abilities{
		c:         c,
		body:      container.NewVBox(),
		abilities: container.NewVBox(),
	}
	e.ExtendBaseWidget(e)

	e.add = widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		e.addRow(&save.Ability{})
	})

	for i, asd := range e.c.AbilitySlotData {
		e.body.Add(widget.NewLabel(fmt.Sprintf("%d", i+1)))
		for _, a := range asd.SlotInfo.Values {
			r := inputs.NewIdEntryWithDataWithHint(&a.AbilityID, finder.Abilities)
			e.body.Add(container.NewPadded(container.NewGridWithColumns(3, r.Label, r.ID)))
		}
	}
	for _, a := range e.c.Abilities {
		e.addRow(a)
	}
	return e
}

func (e *Abilities) addRow(a *save.Ability) {
	var g *fyne.Container
	r := inputs.NewIdEntryWithDataWithHint(&a.AbilityID, finder.Abilities)
	g = container.NewGridWithColumns(4, r.Label, r.ID, widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		e.c.RemoveAbility(a)
		e.abilities.Remove(g)
	}))
	e.abilities.Add(g)
}

func (e *Abilities) CreateRenderer() fyne.WidgetRenderer {
	search := inputs.GetSearches().Abilities
	left := container.NewBorder(
		container.NewGridWithColumns(4, container.NewStack(), widget.NewLabel("Ability ID"), e.add),
		nil, nil, nil,
		container.NewVScroll(e.abilities))
	right := container.NewGridWithColumns(2, search.Fields(), search.Filter())
	return widget.NewSimpleRenderer(
		container.NewBorder(nil, nil,
			left, right,
			container.NewBorder(nil, nil,
				container.NewVScroll(e.body), nil)))
}
