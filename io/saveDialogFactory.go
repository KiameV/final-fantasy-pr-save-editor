package io

import (
	"github.com/aarzilli/nucular"
	"io/ioutil"
	"strings"
)

func SaveInvFile(w *nucular.Window, text []byte) error {
	fn, err := createDialogInv().Save()
	if err != nil {
		return err
	}
	if !strings.Contains(fn, ".ff6inv") {
		fn += ".ff6inv"
	}
	return ioutil.WriteFile(fn, text, 0644)
}
