package save

type (
	Bestiary struct {
		ScenarioFlagsInternal  string `json:"scenarioFlags"`
		MonsterDefeatsInternal string `json:"monsterDefeats"`
		TotalSubjugationCount  int    `json:"totalSubjugationCount"`
	}
	ScenarioFlags struct {
		Target []int `json:"target"`
	}
	MonsterDefeats struct {
		Keys   []int `json:"keys"`
		Values []int `json:"values"`
	}
)

func (b *Bestiary) ScenarioFlags() (v *ScenarioFlags, err error) {
	return UnmarshalOne[ScenarioFlags](b.ScenarioFlagsInternal)
}

func (b *Bestiary) SetScenarioFlags(v *ScenarioFlags) (err error) {
	b.ScenarioFlagsInternal, err = MarshalOne[ScenarioFlags](v)
	return
}

func (b *Bestiary) MonsterDefeats() (v *MonsterDefeats, err error) {
	return UnmarshalOne[MonsterDefeats](b.MonsterDefeatsInternal)
}

func (b *Bestiary) SetMonsterDefeats(v *MonsterDefeats) (err error) {
	b.MonsterDefeatsInternal, err = MarshalOne[MonsterDefeats](v)
	return
}
