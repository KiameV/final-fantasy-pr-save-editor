package core

import (
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/save"
)

type (
	MapData struct {
		Map          *save.MapData
		Player       *save.PlayerEntity
		Gps          *save.GpsData
		OtherParties *[]*save.OtherPartyData
	}
)

func NewMapData(game global.Game, md *save.MapData) (m *MapData, err error) {
	m = &MapData{Map: md}
	if m.Player, err = md.PlayerEntity(); err == nil {
		if m.Gps, err = md.GpsData(); err == nil {
			if game.IsSix() {
				m.OtherParties = md.OtherPartyDataList
			}
		}
	}
	return
}

func (d MapData) ToSave(game global.Game, md *save.MapData) (err error) {
	if err = md.SetGpsData(d.Gps); err == nil {
		if err = md.SetPlayerEntity(d.Player); err != nil {
			if game.IsSix() {
				md.OtherPartyDataList = d.OtherParties
			}
		}
	}
	return
}
