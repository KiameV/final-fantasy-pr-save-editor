package core

import (
	"slices"

	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/save"
)

type (
	Character struct {
		Base               *save.OwnedCharacter
		Parameters         *save.CharacterParameters
		Commands           []int
		Abilities          []*save.Ability
		AbilitySlotData    []*save.AbilitySlotData
		abilitiesLookup    map[int]*save.Ability
		Jobs               []*save.Job
		Equipment          *save.EquipmentList
		AdditionAbilityIds []int
		// SortOrderOwnedAbilityIdsInternal ???
		// AbilityDictionary *save.AbilityDictionary
		SkillLevelTargets *save.SkillLevelTargets
		// LearningAbilities ???
		EquipmentAbilities []int
		EnableCommandsSave bool
		// LearningAbilities  []int
	}
	Characters struct {
		characters []*Character
	}
)

func NewCharacters(game global.Game, in *save.OwnedCharacterList) (c *Characters, err error) {
	var l []*save.OwnedCharacter
	if l, err = in.OwnedCharacters(); err != nil {
		return
	}
	c = &Characters{characters: make([]*Character, len(l))}
	for i, j := range l {
		if c.characters[i], err = NewCharacter(game, j); err != nil {
			return
		}
	}
	return
}

func NewCharacter(game global.Game, in *save.OwnedCharacter) (c *Character, err error) {
	c = &Character{
		Base:            in,
		abilitiesLookup: make(map[int]*save.Ability),
	}

	if c.Parameters, err = in.CharacterParameters(); err != nil {
		return
	}

	var cl *save.CommandList
	if cl, err = in.CommandList(); err != nil {
		return
	}
	c.Commands = cl.Target
	if c.Abilities, err = in.Abilities(); err != nil {
		return
	}
	if c.AbilitySlotData, err = in.AbilitySlotData(); err != nil {
		return
	}
	if c.AdditionAbilityIds, err = c.Base.AdditionOrderOwnedAbilityIds(); err != nil {
		return
	}
	// if c.AbilityDictionary, err = in.AbilityDictionary(); err != nil {
	// 	return
	// }
	// if c.LearningAbilities, err = in.LearningAbilities(); err != nil {
	// 	return
	// }
	if c.SkillLevelTargets, err = in.SkillLevelTargets(); err != nil {
		return
	}
	if c.Equipment, err = in.Equipment(); err != nil {
		return
	}
	if c.Jobs, err = in.Jobs(); err != nil {
		return
	}
	for _, ability := range c.Abilities {
		c.abilitiesLookup[ability.AbilityID] = ability
	}
	return
}

func (c *Characters) All() []*Character {
	return c.characters
}

func (c *Characters) Names() []string {
	names := make([]string, 0, len(c.characters))
	for _, character := range c.characters {
		if character != nil {
			names = append(names, character.Base.Name)
		}
	}
	slices.Sort(names)
	return names
}

func (c *Characters) GetByID(id int) (character *Character, found bool) {
	for _, character = range c.characters {
		if character != nil && character.Base.ID == id {
			return character, true
		}
	}
	return nil, false
}

func (c *Characters) GetByName(name string) (character *Character, found bool) {
	for _, character = range c.characters {
		if character != nil && character.Base.Name == name {
			return character, true
		}
	}
	return nil, false
}

func (c *Character) AddAbility(ability *save.Ability) {
	if _, ok := c.abilitiesLookup[ability.AbilityID]; !ok {
		c.Abilities = append(c.Abilities, ability)
	}
}

func (c *Character) RemoveAbility(ability *save.Ability) {
	for i, a := range c.Abilities {
		if a.AbilityID == ability.AbilityID {
			c.RemoveAbilityIndex(i)
			break
		}
	}
}

func (c *Character) RemoveAbilityIndex(index int) {
	a := c.Abilities[index]
	delete(c.abilitiesLookup, a.AbilityID)
	c.Abilities = append(c.Abilities[:index], c.Abilities[index+1:]...)
}

func (c *Character) AbilitiesToSave() []*save.Ability {
	a := make([]*save.Ability, 0, len(c.Abilities))
	for _, i := range c.Abilities {
		if i.AbilityID != 0 {
			a = append(a, i)
		}
	}
	return a
}

func (c *Character) HasAbility(id int) bool {
	_, ok := c.abilitiesLookup[id]
	return ok
}

func (c *Character) GetAbility(id int) (a *save.Ability, ok bool) {
	a, ok = c.abilitiesLookup[id]
	return
}

func (c *Character) ToSave() (s string, err error) {
	if err = c.Base.SetCharacterParameters(c.Parameters); err != nil {
		return
	}
	if c.EnableCommandsSave {
		if err = c.Base.SetCommandList(c.Commands); err != nil {
			return
		}
	}
	if err = c.Base.SetAbilityList(c.Abilities); err != nil {
		return
	}
	if err = c.Base.SetAbilitySlotData(c.AbilitySlotData); err != nil {
		return
	}
	if err = c.Base.SetAdditionOrderOwnedAbilityIds(c.AdditionAbilityIds); err != nil {
		return
	}
	// if err = c.Base.SetAbilityDictionary(c.AbilityDictionary); err != nil {
	// 	return
	// }
	// if err = c.Base.SetLearningAbilities(c.LearningAbilities); err != nil {
	// 	return
	// }
	if err = c.Base.SetJobs(c.Jobs); err != nil {
		return
	}
	if err = c.Base.SetSkillLevelTargets(c.SkillLevelTargets); err != nil {
		return
	}
	if err = c.Base.SetEquipment(c.Equipment); err != nil {
		return
	}
	return save.MarshalOne(c.Base)
}
