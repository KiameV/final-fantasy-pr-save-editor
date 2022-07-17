package io

import (
	"encoding/json"
	"os"
	"path/filepath"
	"pr_save_editor/global"
)

const file = "prSaveEditor.config"

var config ConfigData

type ConfigData struct {
	WindowX  int    `json:"width"`
	WindowY  int    `json:"height""`
	SaveDir1 string `json:"dir1"`
	SaveDir2 string `json:"dir2"`
	SaveDir3 string `json:"dir3"`
	SaveDir4 string `json:"dir4"`
	SaveDir5 string `json:"dir5"`
	SaveDir6 string `json:"dir6"`
}

func GetConfig() *ConfigData {
	return &config
}

func init() {
	if b, err := os.ReadFile(filepath.Join(global.PWD, file)); err == nil {
		_ = json.Unmarshal(b, &config)
	}
}

func SaveConfig() {
	if f, e1 := os.Create(filepath.Join(global.PWD, file)); e1 == nil {
		if config.WindowX == 0 {
			config.WindowX = global.WindowWidth
		}
		if config.WindowY == 0 {
			config.WindowY = global.WindowHeight
		}
		b, err := json.Marshal(&config)
		if err == nil {
			os.WriteFile(filepath.Join(global.PWD, file), b, 755)
		}
		_, _ = f.Write(b)
	}
}

func (c *ConfigData) GetDir(saveType global.SaveType) (dir string) {
	switch saveType {
	case global.One:
		dir = GetConfig().SaveDir1
	case global.Two:
		dir = GetConfig().SaveDir2
	case global.Three:
		dir = GetConfig().SaveDir3
	case global.Four:
		dir = GetConfig().SaveDir4
	case global.Five:
		dir = GetConfig().SaveDir5
	case global.Six:
		dir = GetConfig().SaveDir6
	}
	if dir == "" {
		dir = "."
	}
	return
}

func (c *ConfigData) SetDir(dir string, saveType global.SaveType) {
	switch saveType {
	case global.One:
		GetConfig().SaveDir1 = dir
	case global.Two:
		GetConfig().SaveDir2 = dir
	case global.Three:
		GetConfig().SaveDir3 = dir
	case global.Four:
		GetConfig().SaveDir4 = dir
	case global.Five:
		GetConfig().SaveDir5 = dir
	case global.Six:
		GetConfig().SaveDir6 = dir
	}
}
