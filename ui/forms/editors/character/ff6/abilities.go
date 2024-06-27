package ff6

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/core/costs/ff6"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Abilities struct {
		widget.BaseWidget
		c           *core.Character
		magicInputs []*inputs.IdEntry
		abilities   []*check
	}
	check struct {
		*widget.Check
		id int
	}
)

func NewAbilities(c *core.Character) *Abilities {
	e := &Abilities{
		c:           c,
		magicInputs: make([]*inputs.IdEntry, len(ff6.MagicSorted)),
	}
	e.ExtendBaseWidget(e)

	e.createMagics()

	switch c.Base.JobID {
	case 3: // Cyan
		e.createAbilities(ff6.Bushidos)
	case 6: // Sabin
		e.createAbilities(ff6.Blitzes)
	case 8: // Stago
		e.createAbilities(ff6.Lores)
	case 11: // Mog
		e.createAbilities(ff6.Dances)
	case 12: // Gau
		e.createAbilities(ff6.Rages)
	}
	return e
}

func (e *Abilities) createMagics() {
	for i, m := range ff6.MagicSorted {
		a, ok := e.c.GetAbility(m.Value)
		if !ok {
			a = &save.Ability{AbilityID: m.Value}
			e.c.AddAbility(a)
		}
		e.magicInputs[i] = inputs.NewLabeledIntEntryWithData(m.Name, &a.SkillLevel)
	}
}

func (e *Abilities) createAbilities(a []models.NameValue) {
	for _, i := range a {
		var c *check
		c = &check{
			Check: widget.NewCheck(i.Name, func(checked bool) {
				e.toggleAbility(c, checked)
			}),
			id: i.Value,
		}
		c.SetChecked(e.knows(i.Value))
		e.abilities = append(e.abilities, c)
	}
}

func (e *Abilities) CreateRenderer() fyne.WidgetRenderer {
	ml := container.NewVBox()
	mm := container.NewVBox()
	mr := container.NewVBox()
	for i, j := range e.magicInputs {
		c := container.NewGridWithColumns(2, j.Label, j.ID)
		if k := i % 3; k == 0 {
			ml.Add(c)
		} else if k == 1 {
			mm.Add(c)
		} else {
			mr.Add(c)
		}
	}

	left := container.NewBorder(
		container.NewHBox(
			widget.NewButton("All", func() { e.toggleAllMagic(true) }),
			widget.NewButton("None", func() { e.toggleAllMagic(false) })),
		nil, nil, nil,
		container.NewVScroll(container.NewGridWithColumns(3, ml, mm, mr)))

	var right *fyne.Container
	if len(e.abilities) == 0 {
		right = container.NewStack()
	} else {
		a := container.NewVBox()
		for _, i := range e.abilities {
			a.Add(i.Check)
		}
		right = container.NewBorder(
			container.NewHBox(
				widget.NewButton("All", func() { e.toggleAllAbilities(true) }),
				widget.NewButton("None", func() { e.toggleAllAbilities(false) })),
			nil, nil, nil,
			container.NewVScroll(a))
	}

	return widget.NewSimpleRenderer(container.NewBorder(nil, nil, left, right))
}

func (e *Abilities) toggleAllMagic(max bool) {
	v := "0"
	if max {
		v = "100"
	}
	for _, i := range e.magicInputs {
		i.ID.SetText(v)
	}
}

func (e *Abilities) toggleAllAbilities(checked bool) {
	for _, i := range e.abilities {
		e.toggleAbility(i, checked)
	}
}

func (e *Abilities) toggleAbility(c *check, checked bool) {
	c.SetChecked(checked)
	ab, ok := e.c.GetAbility(c.id)
	if checked {
		if !ok {
			e.c.AddAbility(&save.Ability{
				AbilityID:  c.id,
				ContentID:  c.id + 330,
				SkillLevel: 100,
			})
		} else {
			ab.SkillLevel = 100
		}
	} else {
		if !ok {
			e.c.AddAbility(&save.Ability{
				AbilityID:  c.id,
				ContentID:  c.id + 330,
				SkillLevel: 0,
			})
		} else {
			ab.SkillLevel = 0
		}
	}
}

func (e *Abilities) knows(id int) bool {
	if a, ok := e.c.GetAbility(id); ok {
		return a.SkillLevel >= 100
	}
	return false
}
