package core

import (
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/save"
)

type (
	MapData struct {
		Map    *save.MapData
		Player *save.PlayerEntity
		Gps    *save.GpsData
	}
)

func NewMapData(game global.Game, md *save.MapData) (m *MapData, err error) {
	m = &MapData{Map: md}
	if m.Player, err = md.PlayerEntity(); err == nil {
		m.Gps, err = md.GpsData()
	}
	return
}

func (d MapData) ToSave(md *save.MapData) (err error) {
	if err = md.SetGpsData(d.Gps); err == nil {
		err = md.SetPlayerEntity(d.Player)
	}
	return
}
