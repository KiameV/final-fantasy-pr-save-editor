package character

import (
	"github.com/aarzilli/nucular"
	"math"
	"pr_save_editor/models"
)

type statsUI struct {
	name              nucular.TextEditor
	selectedCharacter int
}

func newStatsUI() widget {
	u := &statsUI{}

	u.name.Flags = nucular.EditField
	u.name.SingleLine = true
	u.name.Text([]rune{})
	u.name.Maxlen = math.MaxInt
	return u
}

func (u *statsUI) Draw(w *nucular.Window) {
	if len(models.Characters) == 0 {
		return
	}
	if u.selectedCharacter >= len(models.Characters) {
		u.name.SelectAll()
		u.name.DeleteSelection()
		u.selectedCharacter = 0
	}

	c := character

	w.Row(24).Static(50, 200)
	w.Label("Name:", "LC")
	u.name.Edit(w)

	if c.Name != string(u.name.Buffer) {
		n := string(u.name.Buffer)
		models.UpdateCharacterName(character.Name, n)
		character.Name = n
	}

	w.Row(24).Static(200, 200)
	w.PropertyInt("Level", 0, &c.Level, 99, 1, 0)
	w.PropertyInt("Exp", 0, &c.Level, math.MaxUint32, 1000, 0)

	w.Row(24).Static(200, 200)
	w.PropertyInt("Current HP", 0, &c.CurrentHP, 9999, 100, 0)
	w.PropertyInt("Current MP", 0, &c.CurrentMP, 999, 10, 0)

	w.Row(4).Static()

	w.Row(24).Static(400)
	w.Label("Additional Parameters:", "LC")

	w.Row(24).Static(200, 200)
	w.PropertyInt("HP", 0, &c.MaxHP, 9999, 100, 0)
	w.PropertyInt("MP", 0, &c.MaxMP, 999, 10, 0)

	w.Row(24).Static(200, 200)
	w.PropertyInt("Power", 0, &c.Power, 255, 1, 0)
	w.PropertyInt("Vitality", 0, &c.Vitality, 255, 1, 0)

	w.Row(24).Static(200, 200)
	w.PropertyInt("Agility", 0, &c.Agility, 255, 1, 0)
	w.PropertyInt("Weight", 0, &c.Weight, 255, 1, 0)

	w.Row(24).Static(200, 200)
	w.PropertyInt("Intelligence", 0, &c.Intelligence, 255, 1, 0)
	w.PropertyInt("Spirit", 0, &c.Spirit, 255, 1, 0)

	w.Row(24).Static(200, 200)
	w.PropertyInt("Magic", 0, &c.Magic, 255, 1, 0)
	w.PropertyInt("Luck", 0, &c.Luck, 255, 1, 0)

	w.Row(24).Static(200, 200)
	w.PropertyInt("Attack", 0, &c.Attack, 255, 1, 0)
	w.PropertyInt("Defense", 0, &c.Defense, 255, 1, 0)

	w.Row(24).Static(200, 200)
	w.PropertyInt("Ability Defence", 0, &c.AbilityDefence, 255, 1, 0)
	w.PropertyInt("Ability EvasionRate", 0, &c.AbilityEvasionRate, 255, 1, 0)

	w.Row(24).Static(200, 200)
	w.PropertyInt("Accuracy Rate", 0, &c.AccuracyRate, 255, 1, 0)
	w.PropertyInt("Evasion Rate", 0, &c.EvasionRate, 255, 1, 0)

	w.Row(24).Static(200, 200)
	w.PropertyInt("Ability Disturbed Rate", 0, &c.AbilityDisturbedRate, 255, 1, 0)
	w.PropertyInt("Critical Rate", 0, &c.CriticalRate, 255, 1, 0)

	w.Row(24).Static(200, 200)
	w.PropertyInt("Damage Dirmeter", 0, &c.DamageDirmeter, 255, 1, 0)
	w.PropertyInt("Ability Defense Rate", 0, &c.AbilityDefenseRate, 255, 1, 0)

	w.Row(24).Static(200, 200)
	w.PropertyInt("Accuracy Count", 0, &c.AccuracyCount, 255, 1, 0)
	w.PropertyInt("Evasion Count", 0, &c.EvasionCount, 255, 1, 0)

	w.Row(24).Static(200, 200)
	w.PropertyInt("Defense Count", 0, &c.DefenseCount, 255, 1, 0)
	w.PropertyInt("Magic Defense Count", 0, &c.MagicDefenseCount, 255, 1, 0)
}

func (u *statsUI) Update(character *models.Character) {
	u.name.SelectAll()
	u.name.DeleteSelection()
	u.name.Text([]rune(character.Name))
}
