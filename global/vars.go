package global

import (
	"os"
	"time"
)

const (
	WindowWidth  = 725
	WindowHeight = 800
)

var (
	PWD      string
	FileName string
	FileSlot int
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

func NowToTicks() uint64 {
	return uint64(float64(time.Now().UnixNano())*0.01) + uint64(60*60*24*365*1970*10000000)
}
