package characters

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"pr_save_editor/models"
	custom "pr_save_editor/ui/widgets"
)

type (
	TabItem struct {
		*container.TabItem
		grid      *fyne.Container
		Character *models.Character

		Selection *widget.Select
		Enabled   *widget.Check
		Name      *widget.Entry
		Level     *custom.NumericalEntry
		Exp       *custom.NumericalEntry
		CurrentHp *custom.NumericalEntry
		CurrentMp *custom.NumericalEntry

		AdditionalHP         *custom.NumericalEntry
		AdditionalMP         *custom.NumericalEntry
		Power                *custom.NumericalEntry
		Vitality             *custom.NumericalEntry
		Agility              *custom.NumericalEntry
		Weight               *custom.NumericalEntry
		Intelligence         *custom.NumericalEntry
		Spirit               *custom.NumericalEntry
		Magic                *custom.NumericalEntry
		Luck                 *custom.NumericalEntry
		Attack               *custom.NumericalEntry
		Defence              *custom.NumericalEntry
		AbilityDefence       *custom.NumericalEntry
		AbilityEvasionRate   *custom.NumericalEntry
		AccuracyRate         *custom.NumericalEntry
		EvasionRate          *custom.NumericalEntry
		AbilityDisturbedRate *custom.NumericalEntry
		CriticalRate         *custom.NumericalEntry
		DamageDirmeter       *custom.NumericalEntry
		AbilityDefenceRate   *custom.NumericalEntry
		AccuracyCount        *custom.NumericalEntry
		EvasionCount         *custom.NumericalEntry
		DefenceCount         *custom.NumericalEntry
		MagicDefenseCount    *custom.NumericalEntry
	}
)

var (
	tabItem *TabItem
)

func Get() *TabItem {
	if tabItem == nil {
		grid := container.NewGridWithColumns(2)
		tabItem = &TabItem{
			Character: nil,
			grid:      grid,
			Selection: widget.NewSelect([]string{}, func(s string) { tabItem.onSelectedCharacterChanged(s) }),
			Enabled: widget.NewCheck("Enabled", func(b bool) {

			}),
			Name:                 widget.NewEntry(),
			Level:                custom.NewNumericalEntry(1, 99),
			Exp:                  custom.NewNumericalEntry(),
			CurrentHp:            custom.NewNumericalEntry(0, 9999),
			CurrentMp:            custom.NewNumericalEntry(0, 9999),
			AdditionalHP:         custom.NewNumericalEntry(0, 9999),
			AdditionalMP:         custom.NewNumericalEntry(0, 9999),
			Power:                custom.NewNumericalEntry(0, 255),
			Vitality:             custom.NewNumericalEntry(0, 255),
			Agility:              custom.NewNumericalEntry(0, 255),
			Weight:               custom.NewNumericalEntry(0, 255),
			Intelligence:         custom.NewNumericalEntry(0, 255),
			Spirit:               custom.NewNumericalEntry(0, 255),
			Magic:                custom.NewNumericalEntry(0, 255),
			Luck:                 custom.NewNumericalEntry(0, 255),
			Attack:               custom.NewNumericalEntry(0, 255),
			Defence:              custom.NewNumericalEntry(0, 255),
			AbilityDefence:       custom.NewNumericalEntry(0, 255),
			AbilityEvasionRate:   custom.NewNumericalEntry(0, 255),
			AccuracyRate:         custom.NewNumericalEntry(0, 255),
			EvasionRate:          custom.NewNumericalEntry(0, 255),
			AbilityDisturbedRate: custom.NewNumericalEntry(0, 255),
			CriticalRate:         custom.NewNumericalEntry(0, 255),
			DamageDirmeter:       custom.NewNumericalEntry(0, 255),
			AbilityDefenceRate:   custom.NewNumericalEntry(0, 255),
			AccuracyCount:        custom.NewNumericalEntry(0, 255),
			EvasionCount:         custom.NewNumericalEntry(0, 255),
			DefenceCount:         custom.NewNumericalEntry(0, 255),
			MagicDefenseCount:    custom.NewNumericalEntry(0, 255),
		}
		i := tabItem
		binding.NewString()
		i.TabItem = container.NewTabItem("Characters",
			container.NewBorder(i.Selection, nil, nil, nil, grid))

		grid.Add(i.Enabled)
		grid.Add(layout.NewSpacer())
		grid.Add(newLabelEntry("Name", i.Name))
		grid.Add(layout.NewSpacer())
		grid.Add(newLabelEntry("Level", i.Level))
		grid.Add(newLabelEntry("Exp", i.Exp))
		grid.Add(newLabelEntry("Current HP", i.CurrentHp))
		grid.Add(newLabelEntry("Current MP", i.CurrentMp))
		grid.Add(newLabelEntry("Additional HP", i.AdditionalHP))
		grid.Add(newLabelEntry("Additional MP", i.AdditionalMP))
		grid.Add(newLabelEntry("Power", i.Power))
		grid.Add(newLabelEntry("Vitality", i.Vitality))
		grid.Add(newLabelEntry("Agility", i.Agility))
		grid.Add(newLabelEntry("Weight", i.Weight))
		grid.Add(newLabelEntry("Intelligence", i.Intelligence))
		grid.Add(newLabelEntry("Spirit", i.Spirit))
		grid.Add(newLabelEntry("Magic", i.Magic))
		grid.Add(newLabelEntry("Luck", i.Luck))
		grid.Add(newLabelEntry("Attack", i.Attack))
		grid.Add(newLabelEntry("Defence", i.Defence))
		grid.Add(newLabelEntry("Ability Defence", i.AbilityDefence))
		grid.Add(newLabelEntry("Ability Evasion Rate", i.AbilityEvasionRate))
		grid.Add(newLabelEntry("Accuracy Rate", i.AccuracyRate))
		grid.Add(newLabelEntry("Evasion Rate", i.EvasionRate))
		grid.Add(newLabelEntry("Ability Disturbed Rate", i.AbilityDisturbedRate))
		grid.Add(newLabelEntry("Critical Rate", i.CriticalRate))
		grid.Add(newLabelEntry("Damage", i.DamageDirmeter))
		grid.Add(newLabelEntry("Ability Defence Rate", i.AbilityDefenceRate))
		grid.Add(newLabelEntry("Accuracy Count", i.AccuracyCount))
		grid.Add(newLabelEntry("Evasion Count", i.EvasionCount))
		grid.Add(newLabelEntry("Defence Count", i.DefenceCount))
		grid.Add(newLabelEntry("Magic Defense Count", i.MagicDefenseCount))
		grid.Hide()

		i.Enabled.OnChanged = func(b bool) {
			if i.Character != nil {
				i.Character.IsEnabled = b
			}
		}
		i.Name.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.Name = s
			}
		}
		i.Level.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.Level, _ = strconv.Atoi(s)
			}
		}
		i.Exp.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.Exp, _ = strconv.Atoi(s)
			}
		}
		i.CurrentHp.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.CurrentHP, _ = strconv.Atoi(s)
			}
		}
		i.CurrentMp.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.CurrentMP, _ = strconv.Atoi(s)
			}
		}
		i.AdditionalHP.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.MaxHP, _ = strconv.Atoi(s)
			}
		}
		i.AdditionalMP.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.MaxMP, _ = strconv.Atoi(s)
			}
		}
		i.Power.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.Power, _ = strconv.Atoi(s)
			}
		}
		i.Vitality.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.Vitality, _ = strconv.Atoi(s)
			}
		}
		i.Agility.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.Agility, _ = strconv.Atoi(s)
			}
		}
		i.Weight.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.Weight, _ = strconv.Atoi(s)
			}
		}
		i.Intelligence.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.Intelligence, _ = strconv.Atoi(s)
			}
		}
		i.Spirit.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.Spirit, _ = strconv.Atoi(s)
			}
		}
		i.Magic.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.Magic, _ = strconv.Atoi(s)
			}
		}
		i.Luck.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.Luck, _ = strconv.Atoi(s)
			}
		}
		i.Attack.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.Attack, _ = strconv.Atoi(s)
			}
		}
		i.Defence.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.Defense, _ = strconv.Atoi(s)
			}
		}
		i.AbilityDefence.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.AbilityDefence, _ = strconv.Atoi(s)
			}
		}
		i.AbilityEvasionRate.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.AbilityEvasionRate, _ = strconv.Atoi(s)
			}
		}
		i.AccuracyRate.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.AccuracyRate, _ = strconv.Atoi(s)
			}
		}
		i.EvasionRate.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.EvasionRate, _ = strconv.Atoi(s)
			}
		}
		i.AbilityDisturbedRate.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.AbilityDisturbedRate, _ = strconv.Atoi(s)
			}
		}
		i.CriticalRate.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.CriticalRate, _ = strconv.Atoi(s)
			}
		}
		i.DamageDirmeter.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.DamageDirmeter, _ = strconv.Atoi(s)
			}
		}
		i.AbilityDefenceRate.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.AbilityDefenseRate, _ = strconv.Atoi(s)
			}
		}
		i.AccuracyCount.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.AccuracyCount, _ = strconv.Atoi(s)
			}
		}
		i.EvasionCount.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.EvasionCount, _ = strconv.Atoi(s)
			}
		}
		i.DefenceCount.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.DefenseCount, _ = strconv.Atoi(s)
			}
		}
		i.MagicDefenseCount.OnChanged = func(s string) {
			if i.Character != nil {
				i.Character.MagicDefenseCount, _ = strconv.Atoi(s)
			}
		}
	}
	return tabItem
}

func (i *TabItem) Load() {
	possible := []string{""}
	for _, c := range models.Characters {
		possible = append(possible, c.Name)
	}
	i.Selection.Options = possible
	i.Selection.SetSelected("")
}

func (i *TabItem) onSelectedCharacterChanged(s string) {
	if s == "" {
		return
	}
	i.grid.Show()
	var found bool
	if i.Character, found = models.GetCharacterByName(s); found {
		i.Enabled.SetChecked(i.Character.IsEnabled)
		i.Name.SetText(i.Character.Name)
		i.Level.SetText(strconv.Itoa(i.Character.Level))
		i.Exp.SetText(strconv.Itoa(i.Character.Exp))
		i.CurrentHp.SetText(strconv.Itoa(i.Character.CurrentHP))
		i.CurrentMp.SetText(strconv.Itoa(i.Character.CurrentMP))
		i.AdditionalHP.SetText(strconv.Itoa(i.Character.MaxHP))
		i.AdditionalMP.SetText(strconv.Itoa(i.Character.MaxMP))
		i.Power.SetText(strconv.Itoa(i.Character.Power))
		i.Vitality.SetText(strconv.Itoa(i.Character.Vitality))
		i.Agility.SetText(strconv.Itoa(i.Character.Agility))
		i.Weight.SetText(strconv.Itoa(i.Character.Weight))
		i.Intelligence.SetText(strconv.Itoa(i.Character.Intelligence))
		i.Spirit.SetText(strconv.Itoa(i.Character.Spirit))
		i.Magic.SetText(strconv.Itoa(i.Character.Magic))
		i.Luck.SetText(strconv.Itoa(i.Character.Luck))
		i.Attack.SetText(strconv.Itoa(i.Character.Attack))
		i.Defence.SetText(strconv.Itoa(i.Character.Defense))
		i.AbilityDefence.SetText(strconv.Itoa(i.Character.AbilityDefence))
		i.AbilityEvasionRate.SetText(strconv.Itoa(i.Character.AbilityEvasionRate))
		i.AccuracyRate.SetText(strconv.Itoa(i.Character.AccuracyRate))
		i.CriticalRate.SetText(strconv.Itoa(i.Character.CriticalRate))
		i.DamageDirmeter.SetText(strconv.Itoa(i.Character.DamageDirmeter))
		i.AbilityDefenceRate.SetText(strconv.Itoa(i.Character.AbilityDefenseRate))
		i.AccuracyCount.SetText(strconv.Itoa(i.Character.AccuracyCount))
		i.EvasionCount.SetText(strconv.Itoa(i.Character.EvasionCount))
		i.EvasionRate.SetText(strconv.Itoa(i.Character.EvasionRate))
		i.AbilityDisturbedRate.SetText(strconv.Itoa(i.Character.AbilityDisturbedRate))
		i.DefenceCount.SetText(strconv.Itoa(i.Character.DefenseCount))
		i.MagicDefenseCount.SetText(strconv.Itoa(i.Character.MagicDefenseCount))
	}
}

func newLabelEntry(label string, entry fyne.Widget) *fyne.Container {
	return container.NewGridWithColumns(2, widget.NewLabel(label), entry)
}
