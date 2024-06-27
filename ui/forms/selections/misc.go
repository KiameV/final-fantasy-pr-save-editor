package selections

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Misc struct {
		widget.BaseWidget
		misc *core.Misc
	}
)

func NewMisc(misc *core.Misc) *Misc {
	s := &Misc{misc: misc}
	s.ExtendBaseWidget(s)
	return s
}

func (m *Misc) CreateRenderer() fyne.WidgetRenderer {
	i := container.NewVBox()
	i.Add(m.formatInput("Owned Gil:", &m.misc.OwnedGil))
	if m.misc.TotalGil != nil {
		i.Add(m.formatInput("Total Gil:", m.misc.TotalGil))
	}
	i.Add(m.formatInput("Steps:", &m.misc.Steps))
	i.Add(m.formatInput("Number of Saves:", &m.misc.NumberOfSaves))
	i.Add(m.formatInput("Cursed Shield Fight Count:", &m.misc.CursedShieldFightCount))
	i.Add(m.formatInput("Escape Count:", &m.misc.EscapeCount))
	i.Add(m.formatInput("Battle Count:", &m.misc.BattleCount))
	i.Add(m.formatInput("Opened Chest Count:", &m.misc.OpenedChestCount))
	i.Add(container.NewGridWithColumns(2,
		widget.NewLabel("Is Complete Flag:"),
		widget.NewCheckWithData("", binding.BindBool(&m.misc.IsCompleteFlag))))
	i.Add(inputs.NewLabeledEntry("Play Time:", inputs.NewFloatEntryWithData(&m.misc.PlayTime), 2))
	i.Add(m.formatInput("Monsters Killed Count:", &m.misc.MonstersKilledCount))
	i.Add(m.formatInput("Corps Slot Index:", &m.misc.CorpsSlotIndex))
	if m.misc.TimeStamp != nil {
		i.Add(inputs.NewLabeledEntry("Time Stamp:", widget.NewEntryWithData(binding.BindString(m.misc.TimeStamp)), 2))
	}
	return widget.NewSimpleRenderer(container.NewBorder(nil, nil, container.NewVScroll(i), nil))
}

func (m *Misc) formatInput(label string, i *int) fyne.CanvasObject {
	in := inputs.NewLabeledIntEntryWithData(label, i)
	return container.NewGridWithColumns(2, in.Label, in.ID)
}
