package core

import (
	"strings"
	"time"

	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/save"
)

const (
	timestampFormat = "01/02/2006 15:04:05 PM"
)

type (
	Misc struct {
		OwnedGil               int
		TotalGil               *int
		Steps                  int
		NumberOfSaves          int
		CursedShieldFightCount int
		EscapeCount            int
		BattleCount            int
		OpenedChestCount       int
		IsCompleteFlag         bool
		PlayTime               float64
		OwnedCrystals          []bool
		MonstersKilledCount    int
		CorpsSlotIndex         int
		TimeStamp              *string
	}
)

func NewMisc(data *save.Data, ud *save.UserData, ds *save.DataStorage) (m *Misc) {
	m = &Misc{
		OwnedGil:               ud.OwnedGil,
		TotalGil:               ud.TotalGil,
		Steps:                  ud.Steps,
		CorpsSlotIndex:         ud.CorpsSlotIndex,
		NumberOfSaves:          ud.SaveCompleteCount,
		CursedShieldFightCount: 0,
		EscapeCount:            ud.EscapeCount,
		BattleCount:            ud.BattleCount,
		OpenedChestCount:       ud.OpenChestCount,
		PlayTime:               ud.PlayTime,
		MonstersKilledCount:    ud.MonstersKilledCount,
		IsCompleteFlag:         data.IsCompleteFlag == 1,
	}
	m.OwnedCrystals, _ = ud.OwnedCrystalFlags()
	if len(data.TimeStamp) > 2 && data.TimeStamp[1] == '/' {
		data.TimeStamp = "0" + data.TimeStamp
	}
	if _, err := time.Parse(timestampFormat, data.TimeStamp); err == nil {
		m.TimeStamp = &data.TimeStamp
	}
	if data.Game.IsSix() {
		if len(ds.Global) > 9 {
			m.CursedShieldFightCount = ds.Global[9]
		}
	}
	return
}

func (m *Misc) ToSave(game global.Game, ud *save.UserData, ds *save.DataStorage) {
	ud.OwnedGil = m.OwnedGil
	ud.TotalGil = m.TotalGil
	ud.Steps = m.Steps
	ud.CorpsSlotIndex = m.CorpsSlotIndex
	ud.SaveCompleteCount = m.NumberOfSaves
	ud.EscapeCount = m.EscapeCount
	ud.BattleCount = m.BattleCount
	ud.OpenChestCount = m.OpenedChestCount
	ud.PlayTime = m.PlayTime
	ud.MonstersKilledCount = m.MonstersKilledCount
	_ = ud.SetOwnedCrystalFlags(m.OwnedCrystals)
	if game.IsSix() {
		if len(ds.Global) > 9 {
			ds.Global[9] = m.CursedShieldFightCount
		}
	}
}

func (m *Misc) ToSaveData(data *save.Data) {
	if m.IsCompleteFlag {
		data.IsCompleteFlag = 1
	} else {
		data.IsCompleteFlag = 0
	}
	if m.TimeStamp != nil {
		t := *m.TimeStamp
		if len(t) > 2 && t[1] == '/' {
			t = "0" + t
		}
		if _, err := time.Parse(timestampFormat, t); err == nil {
			data.TimeStamp = t
		}
		data.TimeStamp = strings.TrimLeft(data.TimeStamp, "0")
	}
}
