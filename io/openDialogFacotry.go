package io

import (
	"github.com/aarzilli/nucular"
	"github.com/sqweek/dialog"
	"io/fs"
	"io/ioutil"
	"pr_save_editor/global"
)

func OpenDirAndFileDialog(saveType global.SaveType) (dir string, files []fs.FileInfo, err error) {
	d := dialog.Directory()
	d = d.Title("Select Save Directory")
	d = d.SetStartDir(GetConfig().GetDir(saveType))

	if dir, err = d.Browse(); err == nil {
		GetConfig().SetDir(dir, saveType)
		SaveConfig()
		files, err = ioutil.ReadDir(dir)
	}
	return
}

func OpenInvFileDialog(w *nucular.Window) (data []byte, err error) {
	var fn string
	if fn, err = createInvDialog().Load(); err != nil {
		return
	}
	return ioutil.ReadFile(fn)
}
