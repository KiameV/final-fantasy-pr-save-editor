package main

import (
	"os"
	"strconv"

	"pr_save_editor/ui"
)

const version = "1.0.0"

func main() {
	_ = os.RemoveAll("pr_io")
	defer func() {
		if err := recover(); err != nil {
			var msg string
			switch e := err.(type) {
			case string:
				msg = e
			case error:
				msg = e.Error()
			}
			if msg != "" {
				_ = os.WriteFile("log.txt", []byte(msg), 0644)
			}
		}
	}()

	readScaleFile()

	ui.Show(version)
}

func readScaleFile() {
	if b, err := os.ReadFile("setting"); err == nil {
		if _, err = strconv.ParseFloat(string(b), 64); err == nil {
			_ = os.Setenv("FYNE_SCALE", string(b))
		}
	}
}
