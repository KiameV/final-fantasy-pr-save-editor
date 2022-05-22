package mainMenu

import (
	"github.com/aarzilli/nucular"
	"pr_save_editor/ui"
	"pr_save_editor/ui/character"
	"pr_save_editor/ui/inventory"
	"pr_save_editor/ui/misc"
)

type mainMenu struct {
	uis []ui.UI
}

func NewUI() ui.UI {
	return &mainMenu{
		uis: []ui.UI{
			character.NewUI(),
			inventory.NewUI(),
			misc.NewUI(),
			//party.NewUI(),
		},
	}
}

func (u *mainMenu) Draw(w *nucular.Window) {
	w.Row(5).Static(1)

	for i := 0; i < len(u.uis); i++ {
		b := u.uis[i].Behavior()
		if b == ui.Hide {
			continue
		}

		if w.TreePush(nucular.TreeTab, u.uis[i].Name(), false) {
			u.uis[i].Draw(w)
			w.TreePop()
		}
	}
}

func (u *mainMenu) Refresh() {
	for _, ui := range u.uis {
		ui.Refresh()
	}
}

func (u *mainMenu) Name() string {
	return "Main Menu"
}

func (u *mainMenu) Behavior() ui.Behavior {
	return ui.Show
}
