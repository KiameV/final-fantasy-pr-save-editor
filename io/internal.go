package io

import (
	"fmt"
	"github.com/sqweek/dialog"
	"pr_save_editor/global"
)

func createInvDialog() *dialog.FileBuilder {
	d := dialog.File()
	d = d.SetStartDir(".")
	d = d.Title("Select the Inventory Save File").Filter(getName(), getExt())
	return d
}

func createDialogInv() *dialog.FileBuilder {
	d := dialog.File()
	d = d.SetStartDir(".")
	d = d.Title("Select the Inventory Save File").Filter(getName(), getExt())
	return d
}

func getName() string {
	return fmt.Sprintf("FF%dINV File", global.GetSaveType())
}

func getExt() string {
	return fmt.Sprintf("ff%dinv", global.GetSaveType())
}
