package inventory

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pr_save_editor/models"
)

type (
	TabItem struct {
		*container.TabItem
		inventory *widget.List
		rows []row
		items     *container.Scroll
		weapons   *container.Scroll
		armors    *container.Scroll
		finder    *widget.Entry
		found     *container.Scroll
	}
	row struct {
		fyne.CanvasObject
		Index int
		row *models.Row
	}
)

var (
	tabItem *TabItem
)

func Get() *TabItem {
	if tabItem == nil {
		grid := container.NewGridWithColumns(3)
		tabItem = &TabItem{
			TabItem: container.NewTabItem("Inventory", grid),
		}
		tabItem.inventory = widget.NewList(
				func() int {return 255},
				func() fyne.CanvasObject {
					i := len(tabItem.rows)
					r := row{
						Index: i,
						row: models.GetInventory().Rows[i],
					}
					tabItem.rows = append(tabItem.rows, r)
					return container.NewGridWithColumns(3,
						widget.NewLabel())
				},
				func(id widget.ListItemID, row fyne.CanvasObject) {}),
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
