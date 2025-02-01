package bestiary

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/finder"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Bestiary struct {
		widget.BaseWidget
		monstersDefeated []*inputs.IdCountEntry
		// scenarioFlags         []binding.Bool
		search                *inputs.Search
		killAllMonstersButton *widget.Button
		addRowButton          *widget.Button
		rows                  *fyne.Container
	}
)

func NewCore(b *core.Bestiary, finder finder.Find, search *inputs.Search) *Bestiary {
	e := &Bestiary{
		monstersDefeated: make([]*inputs.IdCountEntry, len(b.MonsterDefeats)),
		// scenarioFlags:         make([]binding.Bool, len(b.ScenarioFlags)),
		search: search,
		rows:   container.NewVBox(),
	}
	e.ExtendBaseWidget(e)
	e.killAllMonstersButton = widget.NewButton("Kill All Monsters", func() {
		killed := make(map[int]int)
		defeats := make([]core.MonsterDefeatedCount, 0, len(*search.Entries))
		for _, d := range e.monstersDefeated {
			killed[d.ID.Int()] = d.Count.Int()
			defeats = append(defeats, core.MonsterDefeatedCount{ID: d.ID.Int(), Count: d.Count.Int()})
		}
		for _, i := range *search.Entries {
			if _, found := killed[i.Value]; !found {
				defeats = append(defeats, core.MonsterDefeatedCount{ID: i.Value, Count: 1})
			}
		}
		b.MonsterDefeats = defeats

		e.monstersDefeated = make([]*inputs.IdCountEntry, len(b.MonsterDefeats))
		for i, d := range b.MonsterDefeats {
			e.monstersDefeated[i] = inputs.NewIdCountEntryWithDataWithHint(&d.ID, &d.Count, finder)
		}
		e.buildRows()
	})
	e.addRowButton = widget.NewButton("+", func() {
		d := core.MonsterDefeatedCount{}
		b.MonsterDefeats = append(b.MonsterDefeats, d)
		e.addRow(inputs.NewIdCountEntryWithDataWithHint(&d.ID, &d.Count, finder))
	})

	for i, d := range b.MonsterDefeats {
		e.monstersDefeated[i] = inputs.NewIdCountEntryWithDataWithHint(&d.ID, &d.Count, finder)
	}
	// for i, f := range b.ScenarioFlags {
	// 	j := binding.NewBool()
	// 	if f == 1 {
	// 		_ = j.Set(true)
	// 	}
	// 	e.scenarioFlags[i] = j
	// }
	return e
}

func (b *Bestiary) CreateRenderer() fyne.WidgetRenderer {
	const columns = 6
	id := widget.NewLabel("Monster ID")
	id.Alignment = fyne.TextAlignCenter
	c := widget.NewLabel("Count")
	c.Alignment = fyne.TextAlignCenter

	headers := container.NewPadded(container.NewGridWithColumns(columns, container.NewStack(), id, c, container.NewStack()))

	b.buildRows()

	top := container.NewGridWithColumns(6)
	top.Add(b.killAllMonstersButton)
	top.Add(container.NewStack())
	top.Add(b.addRowButton)

	return widget.NewSimpleRenderer(
		container.NewBorder(
			top, nil, nil, nil,
			// middle
			container.NewBorder(
				headers, nil, nil, nil,
				container.NewVScroll(b.rows))))
}

func (b *Bestiary) buildRows() {
	b.rows.RemoveAll()
	for _, d := range b.monstersDefeated {
		b.addRow(d)
	}
}

func (b *Bestiary) addRow(d *inputs.IdCountEntry) {
	const columns = 6
	entry := d
	g := container.NewGridWithColumns(columns, entry.Label, entry.ID, entry.Count)
	b.rows.Add(container.NewPadded(g))
}
