package main

import (
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/label"
	"github.com/aarzilli/nucular/rect"
	"github.com/aarzilli/nucular/style"
	"image"
	"image/color"
	"pr_save_editor/browser"
	"pr_save_editor/global"
	"pr_save_editor/io"
	"pr_save_editor/ui"
	"pr_save_editor/ui/file"
	mm "pr_save_editor/ui/mainMenu"
	"strings"
	"time"
)

const version = "0.4.4"

var (
	mainMenu      ui.UI
	status        string
	err           error
	fileSelector  *file.FileSelector
	statusTimer   *time.Timer
	errTextEditor nucular.TextEditor
	fs            *file.FileSelector
)

func main() {
	errTextEditor.Flags = nucular.EditReadOnly | nucular.EditSelectable | nucular.EditSelectable | nucular.EditMultiline
	mainMenu = mm.NewUI()
	var (
		x = io.GetConfig().WindowX
		y = io.GetConfig().WindowY
	)
	if x == 0 || y == 0 {
		x = global.WindowWidth
		y = global.WindowHeight
	}
	wnd := nucular.NewMasterWindowSize(0, "Final Fantasy PR Save Editor - "+version, image.Point{X: x, Y: y}, updateWindow)
	wnd.SetStyle(style.FromTable(customTheme, 1.2))
	wnd.Main()
}

func updateWindow(w *nucular.Window) {
	w.MenubarBegin()
	w.Row(12).Static(100, 100, 300, 200)
	if w := w.Menu(label.TA("Load", "LC"), 100, nil); w != nil {
		ui.DrawError = nil
		w.Row(12).Dynamic(1)
		if w.MenuItem(label.TA("I", "LC")) {
			global.SetShowing(global.LoadPR)
			global.SetSaveType(global.One)
			w.Close()
		} else if w.MenuItem(label.TA("II", "LC")) {
			global.SetShowing(global.LoadPR)
			global.SetSaveType(global.Two)
			w.Close()
		} else if w.MenuItem(label.TA("III", "LC")) {
			global.SetShowing(global.LoadPR)
			global.SetSaveType(global.Three)
			w.Close()
		} else if w.MenuItem(label.TA("IV", "LC")) {
			global.SetShowing(global.LoadPR)
			global.SetSaveType(global.Four)
			w.Close()
		} else if w.MenuItem(label.TA("V", "LC")) {
			global.SetShowing(global.LoadPR)
			global.SetSaveType(global.Five)
			w.Close()
		} else if w.MenuItem(label.TA("VI", "LC")) {
			global.SetShowing(global.LoadPR)
			global.SetSaveType(global.Six)
			w.Close()
		}
	}
	if global.IsShowing(global.ShowPR) {
		ui.DrawError = nil
		if w := w.Menu(label.TA("Save", "LC"), 100, nil); w != nil {
			fileSelector = file.NewFileSelector()
			global.SetShowing(global.SavePR)
			w.Close()
		}
	} else {
		w.Spacing(1)
	}

	if w := w.Menu(label.TA("Check For Update", "LC"), 300, nil); w != nil {
		var hasNewer bool
		var latest string
		if hasNewer, latest, err = browser.CheckForUpdate(version); err != nil {
			popupErr(w, err)
		}
		if hasNewer {
			browser.Update(latest)
		} else {
			status = "version is current"
		}
		w.Close()
	}

	if ui.DrawError != nil {
		popupErr(w, ui.DrawError)
	}

	if status != "" {
		w.Label("Status: "+status, "RC")
		if statusTimer != nil {
			statusTimer.Stop()
		}
		statusTimer = time.AfterFunc(2*time.Second, func() { status = "" })

	} else {
		w.Spacing(1)
	}
	w.MenubarEnd()

	switch global.GetCurrentShowing() {
	case global.LoadPR:
		var loaded bool
		fs = file.NewFileSelector()
		if loaded, err = fs.DrawLoad(w); err != nil {
			if err.Error() == "Cancelled" {
				err = nil
			} else {
				popupErr(w, err)
			}
			global.RollbackShowing()
		} else if loaded {
			global.SetShowing(global.ShowPR)
			fileSelector = nil
			mainMenu.Refresh()
		}
	case global.SavePR:
		var saved bool
		if saved, err = fileSelector.DrawSave(w); err != nil {
			popupErr(w, err)
			fileSelector = nil
		} else if saved {
			global.SetShowing(global.ShowPR)
			fileSelector = nil
		}
	case global.ShowPR:
		mainMenu.Draw(w)
	}
}

func popupErr(w *nucular.Window, err error) {
	msg := err.Error()
	sb := strings.Builder{}
	for i := 0; i < len(msg)-7; i++ {
		c := msg[i]
		if c == '\r' {
			continue
		}
		if c == 'U' || c == 'u' {
			s := msg[i : i+5]
			if s == "Users" || s == "users" {
				i += 10
				sb.WriteString("...")
				continue
			}
		}
		sb.WriteByte(c)
	}
	errTextEditor.Buffer = []rune(sb.String())
	w.Master().PopupOpen("Error", nucular.WindowMovable|nucular.WindowTitle|nucular.WindowDynamic, rect.Rect{X: 20, Y: 100, W: 700, H: 600}, true,
		func(w *nucular.Window) {
			w.Row(300).Dynamic(1)
			errTextEditor.Edit(w)
			w.Row(25).Dynamic(1)
			if w.Button(label.T("OK"), false) {
				w.Close()
			}
		})
}

var customTheme = style.ColorTable{
	ColorText:                  color.RGBA{0, 0, 0, 255},
	ColorWindow:                color.RGBA{255, 255, 255, 255},
	ColorHeader:                color.RGBA{242, 242, 242, 255},
	ColorHeaderFocused:         color.RGBA{0xc3, 0x9a, 0x9a, 255},
	ColorBorder:                color.RGBA{0, 0, 0, 255},
	ColorButton:                color.RGBA{185, 185, 185, 255},
	ColorButtonHover:           color.RGBA{215, 215, 215, 255},
	ColorButtonActive:          color.RGBA{200, 200, 200, 255},
	ColorToggle:                color.RGBA{225, 225, 225, 255},
	ColorToggleHover:           color.RGBA{200, 200, 200, 255},
	ColorToggleCursor:          color.RGBA{30, 30, 30, 255},
	ColorSelect:                color.RGBA{175, 175, 175, 255},
	ColorSelectActive:          color.RGBA{190, 190, 190, 255},
	ColorSlider:                color.RGBA{190, 190, 190, 255},
	ColorSliderCursor:          color.RGBA{215, 215, 215, 255},
	ColorSliderCursorHover:     color.RGBA{235, 235, 235, 255},
	ColorSliderCursorActive:    color.RGBA{225, 225, 225, 255},
	ColorProperty:              color.RGBA{225, 225, 225, 255},
	ColorEdit:                  color.RGBA{245, 245, 245, 255},
	ColorEditCursor:            color.RGBA{0, 0, 0, 255},
	ColorCombo:                 color.RGBA{225, 225, 225, 255},
	ColorChart:                 color.RGBA{160, 160, 160, 255},
	ColorChartColor:            color.RGBA{45, 45, 45, 255},
	ColorChartColorHighlight:   color.RGBA{255, 0, 0, 255},
	ColorScrollbar:             color.RGBA{180, 180, 180, 255},
	ColorScrollbarCursor:       color.RGBA{140, 140, 140, 255},
	ColorScrollbarCursorHover:  color.RGBA{150, 150, 150, 255},
	ColorScrollbarCursorActive: color.RGBA{160, 160, 160, 255},
	ColorTabHeader:             color.RGBA{210, 210, 210, 255},
}
