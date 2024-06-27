package core

import (
	"pixel-remastered-save-editor/save"
)

type (
	Config struct {
		ButtonType              int
		BattleModeIndex         int
		BattleSpeedIndex        int
		BattleMessageSpeedIndex int
		IsCursorMemory          bool
		IsKeepAutoBattle        bool
		MessageSpeed            int
		Brightness              *float64
		MasterVolume            *float64
		BgmVolume               *float64
		SeVolume                *float64
		IsLeftActionIcon        bool
		IsLeftVirtualPad        bool
		IsLeftMenuCommand       bool
		IsLeftBattleCommand     bool
		ReequipIndex            int
		AnalogModeIndex         int
	}
)

func NewConfig(cd *save.ConfigData) *Config {
	return &Config{
		ButtonType:              cd.ButtonType,
		BattleModeIndex:         cd.BattleModeIndex,
		BattleSpeedIndex:        cd.BattleSpeedIndex,
		BattleMessageSpeedIndex: cd.BattleMessageSpeedIndex,
		IsCursorMemory:          cd.IsCursorMemory,
		IsKeepAutoBattle:        cd.IsKeepAutoBattle,
		MessageSpeed:            cd.MessageSpeed,
		Brightness:              cd.Brightness,
		MasterVolume:            cd.MasterVolume,
		BgmVolume:               cd.BgmVolume,
		SeVolume:                cd.SeVolume,
		IsLeftActionIcon:        cd.IsLeftActionIcon,
		IsLeftVirtualPad:        cd.IsLeftVirtualPad,
		IsLeftMenuCommand:       cd.IsLeftMenuCommand,
		IsLeftBattleCommand:     cd.IsLeftBattleCommand,
		ReequipIndex:            cd.ReequipIndex,
		AnalogModeIndex:         cd.AnalogModeIndex,
	}
}

func (c Config) ToSave(cd *save.ConfigData) {
	cd.ButtonType = c.ButtonType
	cd.BattleModeIndex = c.BattleModeIndex
	cd.BattleSpeedIndex = c.BattleSpeedIndex
	cd.BattleMessageSpeedIndex = c.BattleMessageSpeedIndex
	cd.IsCursorMemory = c.IsCursorMemory
	cd.IsKeepAutoBattle = c.IsKeepAutoBattle
	cd.MessageSpeed = c.MessageSpeed
	cd.Brightness = c.Brightness
	cd.MasterVolume = c.MasterVolume
	cd.BgmVolume = c.BgmVolume
	cd.SeVolume = c.SeVolume
	cd.IsLeftActionIcon = c.IsLeftActionIcon
	cd.IsLeftVirtualPad = c.IsLeftVirtualPad
	cd.IsLeftMenuCommand = c.IsLeftMenuCommand
	cd.IsLeftBattleCommand = c.IsLeftBattleCommand
	cd.ReequipIndex = c.ReequipIndex
	cd.AnalogModeIndex = c.AnalogModeIndex
}
