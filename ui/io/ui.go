package io

import (
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	sqweek "github.com/sqweek/dialog"
	"pr_save_editor/global"
	"pr_save_editor/io"
)

type (
	slot struct {
		UUID string
		Name string
	}
)

var (
	slots = []slot{
		{
			UUID: "ookrbATYovG3tEOXIH4HqWnsv8TrUlRWzM8AlCmW2mk=",
			Name: "Slot 1",
		},
		{
			UUID: "vgU2wnuaPje2Or53Iqs8Mp=Al6sdM+GM04Iymv229Ow=",
			Name: "Slot 2",
		},
		{
			UUID: "uhHNR4g5QL5twqCc+IhexaltjtBjJnzzcxh5RBSy4G4=",
			Name: "Slot 3",
		},
		{
			UUID: "fmsBRQ+D6YzdjCbBbl7BQuagHyg=7iX3I=EnhccyGDM=",
			Name: "Slot 4",
		},
		{
			UUID: "NXa+MQ+hiHKlPAHJ6GiVWi2Wk5JR2xQQaQxzhyCbK2E=",
			Name: "Slot 5",
		},
		{
			UUID: "UWtRedIOaeA6ig=8r6DIvxg33X92oMM9P8JBwiag4d0=",
			Name: "Slot 6",
		},
		{
			UUID: "e1gfNt2iCE2I3yucQ8zfXn0ou+P2=lREb2q7Lqm04Gc=",
			Name: "Slot 7",
		},
		{
			UUID: "6Pf6Ky7e4QBPuKH9EFJ1Iu+BUEz0zNrXdaS8866Gcq0=",
			Name: "Slot 8",
		},
		{
			UUID: "9dHjN5+9JJWfJ9xoprXo=ehwoEwJwKRYL1Hlc92UNQk=",
			Name: "Slot 9",
		},
		{
			UUID: "oY6N7KlcC4jscZnfa4ea6Nr=TUSR+I=29kwPNZe2NAo=",
			Name: "Slot 10",
		},
		{
			UUID: "NKQ3ux2pea=DqE=vXPKb8+oix5Lt467opYaG0p0brgU=",
			Name: "Slot 11",
		},
		{
			UUID: "HyhjsKWa=tCVf3TWB3qRy7NyrJbc8orciJCntDpqT=I=",
			Name: "Slot 12",
		},
		{
			UUID: "hl9YCUf633k79xePC9PiKAEOq1ajUcSZkLofQuNw2OM=",
			Name: "Slot 13",
		},
		{
			UUID: "C=ozNkSxgKEoLCgOPLJakAUUhnL820LbGlpMz0irQFI=",
			Name: "Slot 14",
		},
		{
			UUID: "z2837SldCS+oIV8y4w5LrnJK9URKYy1QrnoA9bvCg5o=",
			Name: "Slot 15",
		},
		{
			UUID: "CnvUyfaDeqDg3XbVpVWJOj=sPKcGMCV3dR=xM8Ze5jE=",
			Name: "Slot 16",
		},
		{
			UUID: "eQ9Km3NT1WoE4h0hFD90ggFIZayYxfHkIVntc7akYVo=",
			Name: "Slot 17",
		},
		{
			UUID: "Lnbq+GaFOc4ybPZaCf=llI0arXo06rJL32Eu+mCwsLg=",
			Name: "Slot 18",
		},
		{
			UUID: "9GkO1xc52WAzswcEtJxs963MkuCohOHgYj0Fhio=fPE=",
			Name: "Slot 19",
		},
		{
			UUID: "mkYfUr4Mtg0zUmF=6lw+bxRLnbnBYp9ayg1KgploDpQ=",
			Name: "Slot 20",
		},
		{
			UUID: "7nCxyzTwG31W3Zlg70mo751W8ETH1n+Km0dWOzRU84Y=",
			Name: "Quick Save",
		},
	}
	grid *fyne.Container
)

func Show(label string, onSelected func(file string), st global.SaveType, w fyne.Window) {
	var (
		dir = widget.NewEntry()
		d   dialog.Dialog
	)
	grid = container.NewGridWithColumns(2, widget.NewLabel("Name"), widget.NewLabel("Date"))
	dir.OnChanged = func(s string) {
		refresh(d, s, onSelected, st, w)
		d.Refresh()
	}
	d = dialog.NewCustom(label, "Cancel", container.NewBorder(
		container.NewBorder(nil, nil, nil,
			container.NewHBox(
				widget.NewButtonWithIcon("", theme.SearchIcon(), func() {
					if s, err := sqweek.Directory().Title("Load Directory").Browse(); s != "" && err == nil {
						dir.SetText(s)
					}
				}),
				widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
					refresh(d, dir.Text, onSelected, st, w)
				})), dir), nil, nil, nil, grid), w)
	dir.SetText(io.GetConfig().GetDir(st))
	d.Show()
}

func refresh(d dialog.Dialog, dir string, onSelected func(file string), st global.SaveType, w fyne.Window) {
	grid.RemoveAll()
	if dir == "" {
		return
	}
	var i int
	for _, s := range slots {
		if _, err := os.Stat(filepath.Join(dir, s.UUID)); err == nil {
			grid.Add(container.NewBorder(nil, nil, nil,
				widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
					dialog.ShowConfirm("Delete?", "Delete "+s.Name+"?", func(b bool) {
						if b {
							_ = os.Remove(filepath.Join(dir, s.UUID))
						}
					}, w)
				}),
				widget.NewButton(s.Name, func() {
					onSelected(filepath.Join(dir, s.UUID))
					d.Hide()
				})))
			grid.Add(layout.NewSpacer())
			i++
		}
	}
	if i > 0 {
		io.GetConfig().SetDir(dir, st)
	}
}
