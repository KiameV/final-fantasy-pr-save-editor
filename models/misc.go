package models

type Misc struct {
	GP                     int
	Steps                  int
	NumberOfSaves          int
	SaveCountRollOver      int
	MapXAxis               int
	MapYAxis               int
	AirshipXAxis           int
	AirshipYAxis           int
	IsAirshipVisible       bool
	CursedShieldFightCount int
	EscapeCount            int
	BattleCount            int
	OpenedChestCount       int
	IsCompleteFlag         bool
	PlayTime               float64
	OwnedCrystals          [4]bool
}

var misc *Misc

func GetMisc() *Misc {
	if misc == nil {
		misc = &Misc{}
	}
	return misc
}
