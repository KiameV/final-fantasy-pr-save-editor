package io

import (
	"github.com/aarzilli/nucular"
	"io/ioutil"
	"strings"
)

func SaveInvFile(w *nucular.Window, text []byte) error {
	fn, err := saveDialogInv()
	if err != nil {
		return err
	}
	ext := getExt()
	if !strings.Contains(fn, ext) {
		fn += ext
	}
	return ioutil.WriteFile(fn, text, 0644)
}
