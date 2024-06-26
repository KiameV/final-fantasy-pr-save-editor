package ff5

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/core/costs/ff5"
	"pixel-remastered-save-editor/models/finder"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Abilities struct {
		widget.BaseWidget
		c *core.Character
	}
)

func NewAbilities(c *core.Character) *Abilities {
	e := &Abilities{c: c}
	e.ExtendBaseWidget(e)
	return e
}

func (e *Abilities) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewAppTabs(
		container.NewTabItem("Abilities", e.createAbilities(ff5.Abilities)),
		container.NewTabItem("White Magic", e.createMagic(ff5.WhiteMagic)),
		container.NewTabItem("Black Magic", e.createMagic(ff5.BlackMagic)),
		container.NewTabItem("Time Magic", e.createMagic(ff5.TimeMagic)),
		container.NewTabItem("Summon Magic", e.createMagic(ff5.SummonMagic)),
		container.NewTabItem("Spellblade", e.createMagic(ff5.Spellblade)),
		container.NewTabItem("Blue Magic", e.createAbilities(ff5.BlueMagic)),
		container.NewTabItem("Sing", e.createMagic(ff5.Songs))))

}

func (e *Abilities) createAbilities(abilities []models.NameValue) fyne.CanvasObject {
	f := finder.New(abilities)
	search := inputs.NewSearch(abilities)
	rows := container.NewVBox()
	for _, a := range e.c.Abilities {
		if _, ok := f(a.AbilityID); ok {
			e.addRow(rows, a, f, abilities)
		}
	}
	top := container.NewHBox(
		widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
			a := &save.Ability{}
			e.addAbility(a)
			e.addRow(rows, a, f, abilities)
		}))
	return container.NewBorder(
		top, nil,
		container.NewVScroll(rows),
		container.NewGridWithColumns(2, search.Fields(), search.Filter()))
}

func (e *Abilities) addRow(rows *fyne.Container, a *save.Ability, f finder.Find, abilities []models.NameValue) {
	i := inputs.NewIdEntryWithDataWithHint(&a.AbilityID, f)
	var g *fyne.Container
	g = container.NewGridWithColumns(3, i.Label, i.ID, widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		e.removeAbility(a)
		rows.Remove(g)
	}))
	rows.Add(g)
}

func (e *Abilities) createMagic(magics []models.NameValue) fyne.CanvasObject {
	rows := container.NewVBox()
	for _, magic := range magics {
		func(magic models.NameValue) {
			i := widget.NewCheck(magic.Name, func(checked bool) {
				e.toggleMagic(checked, &save.Ability{
					AbilityID: magic.Value,
				})
			})
			i.SetChecked(e.c.HasAbility(magic.Value))
			rows.Add(i)
		}(magic)
	}
	return container.NewVScroll(rows)
}

func (e *Abilities) addAbility(a *save.Ability) {
	e.c.AddAbility(a)
}

func (e *Abilities) removeAbility(a *save.Ability) {
	e.c.RemoveAbility(a)
}

func (e *Abilities) toggleMagic(checked bool, a *save.Ability) {
	if checked {
		e.c.AddAbility(a)
	} else {
		e.c.RemoveAbility(a)
	}
}
