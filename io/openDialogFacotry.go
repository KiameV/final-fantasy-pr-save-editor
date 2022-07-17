package io

import (
	"github.com/aarzilli/nucular"
	"github.com/ncruces/zenity"
	"io/fs"
	"io/ioutil"
	"pr_save_editor/global"
)

func OpenDirAndFileDialog(saveType global.SaveType) (dir string, files []fs.FileInfo, err error) {
	dir, err = zenity.SelectFile(
		zenity.Title("Select Save Directory"),
		zenity.Directory(),
		zenity.Filename(GetConfig().GetDir(saveType)))
	if err == nil {
		GetConfig().SetDir(dir, saveType)
		SaveConfig()
		files, err = ioutil.ReadDir(dir)
	}
	return
}

func OpenInvFileDialog(w *nucular.Window) (data []byte, err error) {
	var fn string
	if fn, err = openInvDialog(); err != nil {
		return
	}
	return ioutil.ReadFile(fn)
}
