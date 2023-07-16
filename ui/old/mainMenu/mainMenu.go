package mainMenu

import (
	"github.com/aarzilli/nucular"
	"pr_save_editor/ui/old"
	"pr_save_editor/ui/old/character"
	"pr_save_editor/ui/old/importantInventory"
	"pr_save_editor/ui/old/inventory"
	"pr_save_editor/ui/old/mapData"
	"pr_save_editor/ui/old/misc"
	"pr_save_editor/ui/old/transportation"
)

type mainMenu struct {
	uis []old.UI
}

func NewUI() old.UI {
	return &mainMenu{
		uis: []old.UI{
			character.NewUI(),
			// party.NewUI(),
			inventory.NewUI(),
			importantInventory.NewUI(),
			mapData.NewUI(),
			transportation.NewUI(),
			misc.NewUI(),
		},
	}
}

func (u *mainMenu) Draw(w *nucular.Window) {
	w.Row(5).Static(1)

	for i := 0; i < len(u.uis); i++ {
		b := u.uis[i].Behavior()
		if b == old.Hide {
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

func (u *mainMenu) Behavior() old.Behavior {
	return old.Show
}
