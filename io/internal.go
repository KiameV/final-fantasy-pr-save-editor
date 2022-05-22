package io

import (
	"github.com/sqweek/dialog"
)

func createInvDialog() *dialog.FileBuilder {
	d := dialog.File()
	d = d.SetStartDir(".")
	d = d.Title("Select the Inventory Save File").Filter("FF6INV File", "ff6inv")
	return d
}

func createDialogInv() *dialog.FileBuilder {
	d := dialog.File()
	d = d.SetStartDir(".")
	d = d.Title("Select the Inventory Save File").Filter("FF6INV File", "ff6inv")
	return d
}
