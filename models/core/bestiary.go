package core

import (
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/save"
)

type (
	Bestiary struct {
		ScenarioFlags  []int
		MonsterDefeats []MonsterDefeatedCount
	}
	MonsterDefeatedCount struct {
		ID    int
		Count int
	}
)

func NewBestiary(game global.Game, bestiary *save.Bestiary) (b *Bestiary, err error) {
	b = &Bestiary{}
	var d *save.MonsterDefeats
	if d, err = bestiary.MonsterDefeats(); err != nil {
		return
	}
	b.MonsterDefeats = make([]MonsterDefeatedCount, len(d.Keys))
	for i := 0; i < len(d.Keys); i++ {
		b.MonsterDefeats[i].ID = d.Keys[i]
		b.MonsterDefeats[i].Count = d.Values[i]
	}

	var f *save.ScenarioFlags
	if f, err = bestiary.ScenarioFlags(); err == nil {
		b.ScenarioFlags = f.Target
	}
	return
}

func (b *Bestiary) Add(row *save.OwnedItems) {
	b.MonsterDefeats = append(b.MonsterDefeats, MonsterDefeatedCount{ID: 0, Count: 0})
}

func (b *Bestiary) ToSave() (bestiary *save.Bestiary, err error) {
	bestiary = &save.Bestiary{}
	
	for i := len(b.MonsterDefeats) - 1; i >= 0; i-- {
		if b.MonsterDefeats[i].ID == 0 || b.MonsterDefeats[i].Count == 0 {
			b.MonsterDefeats = append(b.MonsterDefeats[:i], b.MonsterDefeats[i+1:]...)
		}
	}

	keys := make([]int, len(b.MonsterDefeats))
	values := make([]int, len(b.MonsterDefeats))
	for i, m := range b.MonsterDefeats {
		keys[i] = m.ID
		values[i] = m.Count
	}
	if err = bestiary.SetMonsterDefeats(&save.MonsterDefeats{Keys: keys, Values: values}); err != nil {
		bestiary = nil
		return
	}

	if err = bestiary.SetScenarioFlags(&save.ScenarioFlags{Target: b.ScenarioFlags}); err != nil {
		bestiary = nil
	}

	var i int
	for _, m := range b.MonsterDefeats {
		i += m.Count
	}
	bestiary.TotalSubjugationCount = i
	return
}
