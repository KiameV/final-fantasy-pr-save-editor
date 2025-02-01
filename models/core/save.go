package core

import (
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/save"
)

type (
	Save struct {
		Characters         *Characters
		Party              *Party
		Parties            *Parties
		Inventory          *Inventory
		ImportantInventory *Inventory
		Transportations    *Transportations
		Map                *MapData
		Bestiary           *Bestiary
		Misc               *Misc
		Espers             []int
		// Misc               *Misc
		Data *save.Data
	}
)

func NewSave(data *save.Data) (s *Save, err error) {
	var (
		ud *save.UserData
		md *save.MapData
		ds *save.DataStorage
		cl *save.OwnedCharacterList
		b  *save.Bestiary
		oi []*save.OwnedItems
		so []int
	)
	s = &Save{Data: data}
	if ud, md, ds, b, err = data.Unpack(); err != nil {
		return
	}
	if cl, err = ud.OwnedCharacterList(); err != nil {
		return
	}
	if s.Characters, err = NewCharacters(data.Game, cl); err != nil {
		return
	}
	if s.Party, err = NewParty(data.Game, ud); err != nil {
		return
	}
	if s.Parties, err = NewParties(data.Game, ud); err != nil {
		return
	}
	if oi, err = ud.NormalOwnedItems(); err != nil {
		return
	}
	if so, err = ud.NormalOwnedItemSortIDList(); err != nil {
		return
	}
	s.Inventory = NewInventory(data.Game, oi, so)
	if oi, err = ud.ImportantOwnedItems(); err != nil {
		return
	}
	s.ImportantInventory = NewInventory(data.Game, oi, nil)
	if s.Transportations, err = NewTransportations(data.Game, ud); err != nil {
		return
	}
	if s.Map, err = NewMapData(data.Game, md); err != nil {
		return
	}
	if b != nil {
		if s.Bestiary, err = NewBestiary(data.Game, b); err != nil {
			s.Bestiary = nil
			err = nil
		}
	}
	s.Misc = NewMisc(data, ud, ds)
	if data.Game.IsSix() {
		if s.Espers, err = ud.OwnedMagicStones(); err != nil {
			return
		}
	}
	return s, err
}

func (s *Save) ToSave(game global.Game, slot int) (d *save.Data, err error) {
	if err = s.preSave(); err != nil {
		return
	}
	var (
		ud  *save.UserData
		md  *save.MapData
		ds  *save.DataStorage
		all = s.Characters.All()
		b   *save.Bestiary
		cl  = &save.OwnedCharacterList{Target: make([]string, len(all))}
		oi  []*save.OwnedItems
		so  []int
	)
	if ud, md, ds, b, err = s.Data.Unpack(); err != nil {
		return
	}

	for i, c := range all {
		if cl.Target[i], err = c.ToSave(); err != nil {
			return
		}
	}
	if err = ud.SetOwnedCharacterList(cl); err != nil {
		return
	}
	if err = s.Party.ToSave(ud); err != nil {
		return
	}
	if err = s.Parties.ToSave(ud); err != nil {
		return
	}
	if oi, so, err = s.Inventory.ToSave(); err != nil {
		return
	}
	if err = ud.SetNormalOwnedItems(oi); err != nil {
		return
	}
	if err = ud.SetNormalOwnedItemSortIDList(so); err != nil {
		return
	}
	if oi, _, err = s.ImportantInventory.ToSave(); err != nil {
		return
	}
	if err = ud.SetImportantOwnedItems(oi); err != nil {
		return
	}
	if err = s.Transportations.ToSave(ud); err != nil {
		return
	}
	if err = s.Map.ToSave(game, md); err != nil {
		return
	}
	if s.Bestiary != nil {
		b, _ = s.Bestiary.ToSave()
	}
	s.Misc.ToSave(game, ud, ds)
	if s.Game().IsSix() {
		if err = ud.SetOwnedMagicStones(s.Espers); err != nil {
			return
		}
	}
	if d, err = s.Data.Pack(slot, ud, md, ds, b); err == nil {
		s.Misc.ToSaveData(d)
	}
	return
}

func (s *Save) preSave() (err error) {
	game := s.Game()

	for _, c := range s.Characters.All() {
		if !c.Base.IsEnableCorps {
			continue
		}
		if game == global.One {
			err = s.preCharacterSaveFF1(c)
		} else if game == global.Two {
			err = s.preCharacterSaveFF2(c)
		} else if game == global.Three {
			err = s.preCharacterSaveFF3(c)
		} else if game == global.Four {
			err = s.preCharacterSaveFF4(c)
		} else if game == global.Five {
			err = s.preCharacterSaveFF5(c)
		} else { // six
			err = s.preCharacterSaveFF6(c)
		}
		if err != nil {
			return
		}
	}
	return
}

func (s *Save) preCharacterSaveFF1(c *Character) (err error) {
	owned := make(map[int]bool)
	abilities := make(map[int]bool)
	for _, i := range c.AdditionAbilityIds {
		owned[i] = true
	}
	for _, a := range c.Abilities {
		abilities[a.ContentID] = true
	}
	for _, asd := range c.AbilitySlotData {
		for _, sd := range asd.SlotInfo.Values {
			if sd.AbilityID > 0 {
				sd.ContentID = sd.AbilityID + 208
			}
			if sd.ContentID > 0 {
				if _, ok := owned[sd.ContentID]; !ok {
					c.AdditionAbilityIds = append(c.AdditionAbilityIds, sd.ContentID)
				}
				if _, ok := abilities[sd.ContentID]; !ok {
					c.Abilities = append(c.Abilities, &save.Ability{AbilityID: sd.AbilityID, ContentID: sd.ContentID})
				}
			}
		}
	}
	for _, a := range c.Abilities {
		if a.AbilityID > 0 {
			if _, ok := owned[a.ContentID]; !ok {
				c.AdditionAbilityIds = append(c.AdditionAbilityIds, a.ContentID)
			}
		}
	}
	return
}

func (s *Save) preCharacterSaveFF2(c *Character) (err error) {
	for _, a := range c.Abilities {
		if a.ContentID == 0 && a.AbilityID > 0 {
			a.ContentID = a.AbilityID + 207
		}
	}
	return
}

func (s *Save) preCharacterSaveFF3(c *Character) (err error) {
	for _, a := range c.Abilities {
		if a.ContentID == 0 && a.AbilityID > 0 {
			a.ContentID = a.AbilityID + 201
		}
	}
	return
}

func (s *Save) preCharacterSaveFF4(c *Character) (err error) {
	for _, a := range c.Abilities {
		if a.ContentID == 0 && a.AbilityID > 0 {
			a.ContentID = a.AbilityID + 310
		}
	}
	return
}

func (s *Save) preCharacterSaveFF5(c *Character) (err error) {
	for _, a := range c.Abilities {
		if a.ContentID == 0 && a.AbilityID > 0 {
			a.ContentID = a.AbilityID + 266
		}
	}
	return
}

func (s *Save) preCharacterSaveFF6(c *Character) (err error) {
	for _, a := range c.Abilities {
		if a.ContentID == 0 && a.AbilityID > 0 {
			if a.ContentID == 0 && a.AbilityID > 0 {
				a.ContentID = a.AbilityID + 330
			}
		}
	}
	return
}

func (s *Save) Game() global.Game {
	return s.Data.Game
}
