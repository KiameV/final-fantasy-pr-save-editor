package global

import (
	"io/fs"
	"os"
)

const (
	WindowWidth  = 725
	WindowHeight = 800
)

var (
	PWD      string
	DirFiles []fs.FileInfo
	FileName string
	showing  CurrentScreen
	prevShow CurrentScreen
	saveType SaveType
)

type CurrentScreen byte

const (
	Blank CurrentScreen = iota
	LoadPR
	SavePR
	ShowPR
)

func RollbackShowing() CurrentScreen {
	showing = prevShow
	return showing
}

func GetCurrentShowing() CurrentScreen {
	return showing
}

func SetShowing(s CurrentScreen) {
	if showing != prevShow {
		prevShow = showing
	}
	showing = s
}

func IsShowing(s CurrentScreen) bool {
	return s == showing
}

type SaveType byte

const (
	Unspecified SaveType = iota
	One
	Two
	Three
	Four
	Five
	Six
)

func GetSaveType() SaveType {
	return saveType
}

func SetSaveType(st SaveType) {
	saveType = st
}

func init() {
	var err error
	if PWD, err = os.Getwd(); err != nil {
		PWD = "."
	}
}
