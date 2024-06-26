package ff4

import (
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
		c    *core.Character
		rows []*row
		add  *widget.Button
	}
	row struct {
		*inputs.IdEntry
		remove *widget.Button
	}
)

func NewAbilities(c *core.Character) *Abilities {
	e := &Abilities{
		c: c,
	}
	e.ExtendBaseWidget(e)
	e.add = widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		e.addAbility()
	})
	e.populate()
	return e
}

func (e *Abilities) addAbility() {
	a := &save.Ability{}
	e.c.AddAbility(a)
	e.populate()
}

func (e *Abilities) removeAbility(index int) {
	e.c.RemoveAbilityIndex(index)
	e.populate()
}

func (e *Abilities) populate() {
	e.rows = make([]*row, len(e.c.Abilities))
	for i, a := range e.c.Abilities {
		func(i int, a *save.Ability) {
			e.rows[i] = &row{
				IdEntry: inputs.NewIdEntryWithDataWithHint(&a.AbilityID, finder.Abilities),
				remove: widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
					e.removeAbility(i)
				}),
			}
		}(i, a)
	}
}

func (e *Abilities) CreateRenderer() fyne.WidgetRenderer {
	body := container.NewVBox()
	for _, r := range e.rows {
		body.Add(container.NewGridWithColumns(3, r.Label, r.ID, r.remove))
	}
	search := inputs.GetSearches().Abilities
	left := container.NewBorder(
		container.NewGridWithColumns(3, container.NewStack(), widget.NewLabel("Ability ID"), e.add),
		nil, nil, nil,
		container.NewVScroll(body))
	right := container.NewGridWithColumns(2, search.Fields(), search.Filter())
	return widget.NewSimpleRenderer(container.NewBorder(nil, nil, left, right))
}
