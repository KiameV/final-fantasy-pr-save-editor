package io

import (
	"fmt"
	"github.com/ncruces/zenity"
	"pr_save_editor/global"
)

func openInvDialog() (string, error) {
	return zenity.SelectFile(
		zenity.Title("Select the Inventory Save File"),
		zenity.Filename("."),
		zenity.FileFilter{
			Name:     getName(),
			Patterns: []string{getExt()},
		})
}

func saveDialogInv() (string, error) {
	return zenity.SelectFileSave(
		zenity.Title("Select the Inventory Save File"),
		zenity.Filename("."),
		zenity.FileFilter{
			Name:     getName(),
			Patterns: []string{getExt()},
		})
}

func getName() string {
	return fmt.Sprintf("FF%dINV File", global.GetSaveType())
}

func getExt() string {
	return fmt.Sprintf("*.ff%dinv", global.GetSaveType())
}
