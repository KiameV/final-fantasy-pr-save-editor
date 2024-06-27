package editors

import (
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/core/costs/ff6"
)

type (
	Espers struct {
		widget.BaseWidget
		hasEspers map[int]bool
		rows      []*widget.Check
		espers    *[]int
	}
)

func NewEspers(save *core.Save) *Espers {
	e := &Espers{
		hasEspers: make(map[int]bool),
		rows:      make([]*widget.Check, 0, len(ff6.EspersSorted)),
		espers:    &save.Espers,
	}
	e.ExtendBaseWidget(e)

	for _, i := range save.Espers {
		e.hasEspers[i] = true
	}
	for _, i := range ff6.EspersSorted {
		func(esper models.NameValue) {
			c := widget.NewCheck(esper.Name, func(checked bool) {
				e.toggleEsper(esper, checked)
			})
			_, checked := e.hasEspers[esper.Value]
			c.SetChecked(checked)
			e.rows = append(e.rows, c)
		}(i)
	}
	return e
}

func (e *Espers) CreateRenderer() fyne.WidgetRenderer {
	left := container.NewVBox()
	right := container.NewVBox()
	half := len(e.rows) / 2
	for i, c := range e.rows {
		if i <= half {
			left.Add(c)
		} else {
			right.Add(c)
		}
	}
	return widget.NewSimpleRenderer(container.NewBorder(nil, nil,
		container.NewVScroll(container.NewGridWithColumns(2, left, right)), nil))
}

func (e *Espers) toggleEsper(esper models.NameValue, checked bool) {
	if checked {
		if _, checked = e.hasEspers[esper.Value]; !checked {
			e.hasEspers[esper.Value] = true
			*e.espers = append(*e.espers, esper.Value)
			slices.Sort(*e.espers)
		}
	} else {
		if _, checked = e.hasEspers[esper.Value]; checked {
			delete(e.hasEspers, esper.Value)
			for i := 0; i < len(*e.espers); i++ {
				if (*e.espers)[i] == esper.Value {
					*e.espers = append((*e.espers)[:i], (*e.espers)[i+1:]...)
				}
			}
		}
	}
}
