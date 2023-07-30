package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"pr_save_editor/global"
	"pr_save_editor/io/pr"
	"pr_save_editor/ui/io"
)

var (
	m *fyne.MainMenu
)

func Get(onLoaded func(), w fyne.Window) *fyne.MainMenu {
	if m == nil {
		l := fyne.NewMenuItem("Load", func() {})
		l.ChildMenu = fyne.NewMenu("",
			fyne.NewMenuItem("I", func() {
				load(onLoaded, global.One, w)
			}),
			fyne.NewMenuItem("II", func() {
				load(onLoaded, global.Two, w)
			}),
			fyne.NewMenuItem("III", func() {
				load(onLoaded, global.Three, w)
			}),
			fyne.NewMenuItem("IV", func() {
				load(onLoaded, global.Four, w)
			}),
			fyne.NewMenuItem("V", func() {
				load(onLoaded, global.Five, w)
			}),
			fyne.NewMenuItem("VI", func() {
				load(onLoaded, global.Six, w)
			}))
		s := fyne.NewMenuItem("Save", func() {
			save(w)
		})
		sa := fyne.NewMenuItem("Save As...", func() {
			saveAs(w)
		})
		m = fyne.NewMainMenu(
			fyne.NewMenu("File", l, s, sa))
	}
	return m
}

func AddSave() {
	for _, i := range m.Items[0].Items {
		if i.Label == "Save" {
			return
		}
	}
	m.Items[0].Items = append(m.Items[0].Items, fyne.NewMenuItem("Save", func() {

	}))
	m.Refresh()
}

func load(onLoaded func(), st global.SaveType, w fyne.Window) {
	io.Show(io.Load, func(slot int, file string) {
		if err := pr.NewPR().Load(file); err != nil {
			dialog.ShowError(err, w)
		} else {
			onLoaded()
			global.FileSlot = slot
			global.FileName = file
			global.SetSaveType(st)
		}
	}, st, w)
}

func save(w fyne.Window) {
	if global.FileName != "" && global.FileSlot > 0 {
		if err := pr.NewPR().Save(global.FileSlot, global.FileName); err != nil {
			dialog.ShowError(err, w)
		}
	} else {
		saveAs(w)
	}
}

func saveAs(w fyne.Window) {
	io.Show(io.Save, func(slot int, file string) {
		if err := pr.NewPR().Save(slot, file); err != nil {
			dialog.ShowError(err, w)
		} else {
			global.FileSlot = slot
			global.FileName = file
		}
	}, global.GetSaveType(), w)
}
