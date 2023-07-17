package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"pr_save_editor/ui/characters"
	"pr_save_editor/ui/menu"
)

var (
	application fyne.App
	window      fyne.Window
)

func Show(version string) {
	application = app.New()
	window = application.NewWindow("Save Editor " + version)
	window.Resize(fyne.Size{
		Width:  750,
		Height: 750,
	})
	window.SetMaster()

	window.SetMainMenu(menu.Get(onLoaded, window))
	window.SetContent(container.NewMax(
		container.NewAppTabs(
			characters.Get().TabItem)))
	window.ShowAndRun()
}

func onLoaded() {
	characters.Get().Load()
}
