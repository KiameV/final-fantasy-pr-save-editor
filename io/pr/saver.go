package pr

import (
	"encoding/json"
	"errors"
	"fmt"
	jo "gitlab.com/c0b/go-ordered-json"
	"os"
	"os/exec"
	"path/filepath"
	"pr_save_editor/global"
	"pr_save_editor/io"
	"pr_save_editor/models"
	"strings"
)

func (p *PR) Save(slot int, fileName string) (err error) {
	var (
		toFile = filepath.Join(io.GetConfig().GetDir(global.GetSaveType()), fileName)
		temp   = filepath.Join(global.PWD, "temp")
		cmd    = exec.Command("python", "./io/python/io.py", "obfuscateFile", toFile, temp)
		//needed   = make(map[int]int)
		slTarget = jo.NewOrderedMap()
	)

	if err = p.saveCharacters(); err != nil {
		return
	}
	if err = p.saveInventory(); err != nil {
		return
	}
	if err = p.saveMiscStats(); err != nil {
		return
	}
	if models.GetParty().Enabled {
		if err = p.saveParty(); err != nil {
			return
		}
	}

	iSlice := make([]interface{}, 0, len(p.Characters))
	for _, c := range p.Characters {
		if c != nil {
			var k []byte
			if k, err = c.MarshalJSON(); err != nil {
				return
			}
			s := string(k)
			iSlice = append(iSlice, s)
		}
	}

	if err = p.unmarshalFrom(p.UserData, OwnedCharacterList, slTarget); err != nil {
		return
	}
	slTarget.Set(targetKey, iSlice)
	if err = p.marshalTo(p.UserData, OwnedCharacterList, slTarget); err != nil {
		return
	}

	if err = p.marshalTo(p.Base, UserData, p.UserData); err != nil {
		return
	}

	if err = p.marshalTo(p.Base, MapData, p.MapData); err != nil {
		return
	}

	if err = p.setValue(p.Base, "id", slot); err != nil {
		return
	}

	var data []byte
	if data, err = json.Marshal(p.Base); err != nil {
		return
	}

	if _, err = os.Stat(temp); errors.Is(err, os.ErrNotExist) {
		if _, err = os.Create(temp); err != nil {
			return fmt.Errorf("failed to create save file %s: %v", toFile, err)
		}
	}
	defer os.Remove(temp)

	/*/ TODO Debug
	if _, err = os.Stat("saved.json"); errors.Is(err, os.ErrNotExist) {
		if _, err = os.Create("saved.json"); err != nil {
			return fmt.Errorf("failed to create save file %s: %v", toFile, err)
		}
	}
	s := string(p.data)
	s = strings.ReplaceAll(s, `\`, ``)
	s = strings.ReplaceAll(s, `"{`, `{`)
	s = strings.ReplaceAll(s, `}"`, `}`)
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, []byte(s), "", "\t")
	if err != nil {
		err = ioutil.WriteFile("saved.json", prettyJSON.Bytes(), 0644)
	}
	// TODO END /*/

	if len(p.names) > 0 {
		data = p.revertUnicodeNames(data)
	}

	if err = os.WriteFile(temp, data, 0644); err != nil {
		return fmt.Errorf("failed to create temp file %s: %v", toFile, err)
	}

	if _, err = os.Stat(toFile); errors.Is(err, os.ErrNotExist) {
		if _, err = os.Create(toFile); err != nil {
			return fmt.Errorf("failed to create save file %s: %v", toFile, err)
		}
	}

	var out []byte
	out, err = cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			err = errors.New(string(ee.Stderr))
		} else {
			err = fmt.Errorf("%s: %v", string(out), err)
		}
	}
	return
}

func (p *PR) saveCharacters() (err error) {
	for _, d := range p.Characters {
		if d == nil {
			continue
		}

		var id int
		if id, err = p.getInt(d, ID); err != nil {
			return
		}

		c, found := models.GetCharacter(id)
		if !found {
			continue
		}

		if err = p.setValue(d, JobID, c.JobID); err != nil {
			return
		}

		if err = p.setValue(d, Name, c.Name); err != nil {
			return
		}

		if err = p.setValue(d, IsEnableCorps, c.IsEnabled); err != nil {
			return
		}

		params := jo.NewOrderedMap()
		if err = p.unmarshalFrom(d, Parameter, params); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalLevel, c.Level); err != nil {
			return
		}

		if err = p.setValue(params, CurrentHP, c.CurrentHP); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalMaxHp, c.MaxHP); err != nil {
			return
		}

		if err = p.setValue(params, CurrentMP, c.CurrentMP); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalMaxMp, c.MaxMP); err != nil {
			return
		}

		if err = p.setValue(d, CurrentExp, c.Exp); err != nil {
			return
		}

		// TODO Status

		if err = p.setValue(params, AdditionalPower, c.Power); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalVitality, c.Vitality); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalAgility, c.Agility); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalWeight, c.Weight); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalIntelligence, c.Intelligence); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalSpirit, c.Spirit); err != nil {
			return
		}

		if err = p.setValue(params, AdditionAttack, c.Attack); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalDefence, c.Defense); err != nil {
			return
		}

		if err = p.setValue(params, AdditionAbilityDefense, c.AbilityDefence); err != nil {
			return
		}

		if err = p.setValue(params, AdditionAbilityEvasionRate, c.AbilityEvasionRate); err != nil {
			return
		}

		if err = p.setValue(params, AdditionMagic, c.Magic); err != nil {
			return
		}

		if err = p.setValue(params, AdditionLuck, c.Luck); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalAccuracyRate, c.AccuracyRate); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalEvasionRate, c.EvasionRate); err != nil {
			return
		}

		if err = p.setValue(params, AdditionAbilityDistrurbedRate, c.AbilityDisturbedRate); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalCriticalRate, c.CriticalRate); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalDamageDirmeter, c.DamageDirmeter); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalAbilityDefenseRate, c.AbilityDefenseRate); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalAccuracyCount, c.AccuracyCount); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalEvasionCount, c.EvasionCount); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalDefenceCount, c.DefenseCount); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalMagicDef, c.MagicDefenseCount); err != nil {
			return
		}

		if err = p.marshalTo(d, Parameter, params); err != nil {
			return
		}
	}
	return
}

func (p *PR) addToNeeded(needed *map[int]int, id int) {
	if count, found := (*needed)[id]; !found {
		(*needed)[id] = 1
	} else {
		(*needed)[id] = count + 1
	}
}

type partyMember struct {
	ID          int `json:"id"`
	CharacterID int `json:"characterId"`
}

func (p *PR) saveParty() (err error) {
	var (
		party   = models.GetParty()
		partyID = p.getPartyID()
		b       []byte
		sl      = make([]interface{}, 4)
	)
	for i, m := range party.Members {
		pm := partyMember{
			ID:          partyID,
			CharacterID: m.CharacterID,
		}
		if b, err = json.Marshal(&pm); err != nil {
			return
		}
		sl[i] = string(b)
	}
	return p.setTarget(p.UserData, CorpsList, sl)
}

func (p *PR) getPartyID() int {
	i, err := p.getFromTarget(p.UserData, CorpsList)
	if err != nil {
		return 1
	}
	var m partyMember
	for _, c := range i.([]interface{}) {
		if err = json.Unmarshal([]byte(c.(string)), &m); err != nil {
			return 1
		}
	}
	return m.ID
}

func (p *PR) saveInventory() (err error) {
	var (
		rows             = models.GetInventory().GetRowsForPrSave()
		sl               = make([]interface{}, 0, len(rows))
		b                []byte
		slTarget         = jo.NewOrderedMap()
		found            = make(map[int]bool)
		removeDuplicates = models.GetInventory().RemoveDuplicates
		addedCountLookup = make(map[int]int)
	)

	for _, r := range rows {
		if removeDuplicates {
			if _, f := found[r.ItemID]; f {
				continue
			}
			found[r.ItemID] = true
		}
		// Skip Empty rows
		if r.ItemID == 0 || r.Count == 0 {
			continue
		}
		if b, err = json.Marshal(r); err != nil {
			return
		}
		sl = append(sl, string(b))
	}

	for k, v := range addedCountLookup {
		if b, err = json.Marshal(&models.Row{
			ItemID: k,
			Count:  v,
		}); err != nil {
			return
		}
		sl = append(sl, string(b))
	}

	slTarget.Set(targetKey, sl)
	if err = p.marshalTo(p.UserData, NormalOwnedItemList, slTarget); err != nil {
		return
	}

	if models.GetInventory().ResetSortOrder {
		slTarget = jo.NewOrderedMap()
		slTarget.Set(targetKey, make([]interface{}, 0))
		if err = p.marshalTo(p.UserData, NormalOwnedItemSortIdList, slTarget); err != nil {
			return
		}
	}
	return
}

func (p *PR) saveMiscStats() (err error) {
	misc := models.GetMisc()
	if err = p.setValue(p.UserData, OwnedGil, misc.GP); err != nil {
		return
	}
	if err = p.setValue(p.UserData, Steps, misc.Steps); err != nil {
		return
	}
	if err = p.setValue(p.UserData, EscapeCount, misc.EscapeCount); err != nil {
		return
	}
	if err = p.setValue(p.UserData, BattleCount, misc.BattleCount); err != nil {
		return
	}
	if err = p.setValue(p.UserData, SaveCompleteCount, misc.NumberOfSaves); err != nil {
		return
	}

	if ds, ok := p.Base.GetValue(DataStorage); ok {
		m := jo.NewOrderedMap()
		if err = m.UnmarshalJSON([]byte(ds.(string))); err != nil {
			return
		}
		if err = p.SetIntInSlice(m, "global", misc.CursedShieldFightCount); err != nil {
			return
		}
		var b []byte
		if b, err = m.MarshalJSON(); err != nil {
			return
		}
		p.Base.Set(DataStorage, string(b))
	}

	if err = p.setValue(p.UserData, OpenChestCount, misc.OpenedChestCount); err != nil {
		return
	}
	if err = p.setFlag(p.Base, IsCompleteFlag, misc.IsCompleteFlag); err != nil {
		return
	}
	return p.setValue(p.UserData, PlayTime, misc.PlayTime)
}

func (p *PR) SetIntInSlice(to *jo.OrderedMap, key string, value int) (err error) {
	var (
		i, ok = to.GetValue(key)
		sl    []interface{}
	)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
		return
	}

	if sl, ok = i.([]interface{}); !ok || len(sl) < 9 {
		err = fmt.Errorf("unable to load cursed shield battle count")
		return
	}
	sl[9] = value
	to.Set("global", sl)
	return
}

func (p *PR) setValue(to *jo.OrderedMap, key string, value interface{}) (err error) {
	if !to.Has(key) {
		err = fmt.Errorf("unable to find %s", key)
	}
	to.Set(key, value)
	return
}

func (p *PR) setFlag(to *jo.OrderedMap, key string, value bool) error {
	var i int
	if value {
		i = 1
	}
	return p.setValue(to, key, i)
}

func (p *PR) marshalTo(to *jo.OrderedMap, key string, value *jo.OrderedMap) error {
	if !to.Has(key) {
		return fmt.Errorf("unable to find %s", key)
	}
	if v, err := value.MarshalJSON(); err != nil {
		return err
	} else {
		to.Set(key, string(v))
	}
	return nil
}

func floor0(i int) int {
	if i < 0 {
		return 0
	}
	return i
}

func (p *PR) getInvCount(eq *[]string, counts map[int]int, addedItems *[]int, id int, emptyID int) {
	var i idCount
	if id == 0 {
		i.ContentID = emptyID
		i.Count = counts[emptyID]
	}

	if count, ok := counts[id]; ok {
		i.ContentID = id
		i.Count = count
	} else {
		//*addedItems = append(*addedItems, id)
		i.ContentID = id
		i.Count = 1
	}
	b, _ := json.Marshal(&i)
	*eq = append(*eq, string(b))
}

func (p *PR) setTarget(d *jo.OrderedMap, key string, value []interface{}) (err error) {
	var (
		t = jo.NewOrderedMap()
		b []byte
	)
	if value != nil {
		t.Set(targetKey, value)
	} else {
		t.Set(targetKey, make([]interface{}, 0))
	}
	b, err = t.MarshalJSON()
	return p.setValue(d, key, string(b))
}

func (p *PR) revertUnicodeNames(b []byte) []byte {
	s := string(b)
	for _, r := range p.names {
		s = strings.Replace(s, r.Replaced, r.Original, 1)
	}
	return []byte(s)
	//strconv.Unquote(strings.Replace(strconv.Quote(string(original)), `\\x`, `\x`, -1));
	/*i := 0
	for j := 0; j < len(p.names); j++ {
		original := p.names[j].Original
		replaced := p.names[j].Replaced
		for ; i < len(b)-10; i++ {
			if b[i] == replaced[0] {
				matched := true
				for k := 1; k < len(replaced); k++ {
					if b[i+k] != replaced[k] {
						matched = false
						break
					}
				}
				if matched {
					for k := 0; k < len(replaced); k++ {
						b[i+k] = original[k]
					}
					i += len(replaced)
					break
				}
			}
		}
	}
	return b*/
}
