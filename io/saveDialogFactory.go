package io

import (
	"fmt"
	"github.com/aarzilli/nucular"
	"io/ioutil"
	"pr_save_editor/global"
	"strings"
)

func SaveInvFile(w *nucular.Window, text []byte) error {
	fn, err := createDialogInv().Save()
	if err != nil {
		return err
	}
	ext := fmt.Sprintf(".ff%dinv", global.GetSaveType())
	if !strings.Contains(fn, ext) {
		fn += ext
	}
	return ioutil.WriteFile(fn, text, 0644)
}
