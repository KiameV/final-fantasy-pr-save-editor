package character

import (
	"fmt"

	"github.com/aarzilli/nucular"
	"pr_save_editor/global"
	"pr_save_editor/models"
	"pr_save_editor/ui/old"
	"pr_save_editor/ui/old/file"
)

var character *models.Character

type widget interface {
	Draw(w *nucular.Window)
	Update(character *models.Character)
}

type characterUI struct {
	characterIndex int
	expandAll      bool
	stats          widget
}

func NewUI() old.UI {
	u := &characterUI{
		characterIndex: 0,
		expandAll:      false,
		stats:          newStatsUI(),
	}
	u.Refresh()
	return u
}

func (u *characterUI) Draw(w *nucular.Window) {
	w.Row(18).Static(100, 10, 100)
	if i := w.ComboSimple(models.CharacterNames, u.characterIndex, 12); i != u.characterIndex {
		u.characterIndex = i
		u.refreshWithCharacter(models.Characters[i])
	}
	w.Spacing(1)
	w.CheckboxText("Auto-Expand All", &u.expandAll)

	w.Row(18).Static(300)
	if !character.IsEnabled {
		w.CheckboxText("Disabled (cannot be used in the game)", &character.IsEnabled)
	} else {
		w.CheckboxText("Enabled", &character.IsEnabled)

		// if w.TreePush(nucular.TreeTab, u.makeLabel("Stats"), true) {
		u.stats.Draw(w)
		//	w.TreePop()
		// }
		// if w.TreePush(nucular.TreeTab, u.makeLabel("Status Effects"), u.expandAll) {
		//	u.statusEffects.Draw(w)
		//	w.TreePop()
		// }
	}
}

func (u *characterUI) Refresh() {
	u.characterIndex = 0
	if len(models.Characters) > 0 {
		character = models.Characters[u.characterIndex]
		u.refreshWithCharacter(character)
	}
}

func (u *characterUI) refreshWithCharacter(c *models.Character) {
	character = c
	u.stats.Update(c)
}

func (u *characterUI) Name() string {
	return "Characters"
}

func (u *characterUI) Behavior() old.Behavior {
	return old.Show
}

func (u *characterUI) makeLabel(label string) string {
	name := character.Name
	if global.IsShowing(global.ShowPR) && file.PrIO != nil && file.PrIO.HasUnicodeNames() {
		name = models.Characters[u.characterIndex].Name
	}
	return fmt.Sprintf("%s - %s", label, name)
}
