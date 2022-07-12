package pr

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	jo "gitlab.com/c0b/go-ordered-json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"pr_save_editor/global"
	"pr_save_editor/models"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func (p *PR) Load(fileName string) (err error) {
	var (
		out   []byte
		i     interface{}
		s     string
		names []unicodeNameReplace
	)
	models.GetInventory().Clear()
	models.GetParty().Clear()
	models.Characters = make([]*models.Character, 0, len(p.Characters))
	models.CharacterNames = make([]string, 0, len(p.Characters))

	//p.names = make([][]rune, 0, 40)
	if out, err = p.readFile(fileName); err != nil {
		return
	}
	//ioutil.WriteFile("loaded.json", out, 0644)
	/*for i := 0; i < len(out); i++ {
		if out[i] == '\\' && out[i+1] == 'x' {
			j := string(out[i-20 : i+40])
			print(j)
			i += 50
		}
	}*/
	if s, err = p.getSaveData(string(out)); err != nil {
		return
	}
	//if err == nil {
	//err = ioutil.WriteFile("loaded_pre.json", out, 0644)
	//}
	/*if strings.Contains(s, "\\x") {
		// For foreign langauge, need to double-escape the x
		p.getUnicodeNames(s)
		//for i, n := range p.names {
		//	s = strings.Replace(s, n, fmt.Sprintf(";;;%d;;;", i), 1)
		//}
		s = strings.ReplaceAll(s, "\\x", "x")
	}*/
	s = strings.ReplaceAll(s, `\\r\\n`, "")
	s = p.fixEscapeCharsForLoad(s)
	if strings.Contains(s, "\\x") {
		b := []byte(s)
		if b, err = p.replaceUnicodeNames(b, &names); err != nil {
			return
		}
		s = string(b)
	}
	//s = p.fixFile(s)

	if len(os.Args) >= 2 && os.Args[1] == "print" {
		if _, err = os.Stat("loaded.json"); errors.Is(err, os.ErrNotExist) {
			if _, err = os.Create("loaded.json"); err != nil {
			}
		}
		t := strings.ReplaceAll(s, `\`, ``)
		t = strings.ReplaceAll(t, `"{`, `{`)
		t = strings.ReplaceAll(t, `}"`, `}`)
		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, []byte(t), "", "\t")
		if err == nil {
			err = ioutil.WriteFile("loaded.json", prettyJSON.Bytes(), 0644)
		}
	}

	if err = p.loadBase(s); err != nil {
		return
	}

	//if err = p.Base.UnmarshalJSON([]byte(s)); err != nil {
	//	return
	//}

	if err = p.unmarshalFrom(p.Base, UserData, p.UserData); err != nil {
		return
	}

	if err = p.unmarshalFrom(p.Base, MapData, p.MapData); err != nil {
		return
	}

	if i, err = p.getFromTarget(p.UserData, OwnedCharacterList); err != nil {
		return
	}

	for j, c := range i.([]interface{}) {
		if p.Characters[j] == nil {
			p.Characters[j] = jo.NewOrderedMap()
		}
		s = c.(string)
		if err = p.Characters[j].UnmarshalJSON([]byte(s)); err != nil {
			return
		}
	}

	models.GetParty().Clear()

	if err = p.loadCharacters(); err != nil {
		return
	}
	if err = p.loadParty(); err != nil {
		return
	}
	if err = p.loadMiscStats(); err != nil {
		return
	}
	if err = p.loadInventory(NormalOwnedItemList, models.GetInventory()); err != nil {
		return
	}
	if err = p.loadInventory(importantOwnedItemList, models.GetImportantInventory()); err != nil {
		return
	}
	if err = p.loadMapData(); err != nil {
		return
	}
	if err = p.loadTransportation(); err != nil {
		return
	}

	if len(names) > 0 {
		p.names = names
	}
	return
}

func (p *PR) loadParty() (err error) {
	var (
		party = models.GetParty()
		//id     int
		//name   string
		i      interface{}
		member models.Member
	)
	/*for _, d := range p.Characters {
		if d == nil {
			continue
		}
		if id, err = p.getInt(d, ID); err != nil {
			return
		}
		if name, err = p.getString(d, Name); err != nil {
			return
		}
		party.AddPossibleMember(&models.Member{
			CharacterID: id,
			Name:        name,
		})
	}*/

	if i, err = p.getFromTarget(p.UserData, CorpsList); err != nil {
		return
	}
	for slot, c := range i.([]interface{}) {
		if err = json.Unmarshal([]byte(c.(string)), &member); err != nil {
			return
		}
		party.SetMemberByID(slot, member.CharacterID)
	}
	return
}

func (p *PR) loadCharacters() (err error) {
	party := models.GetParty()
	party.Clear()
	for _, d := range p.Characters {
		if d == nil {
			continue
		}

		c := &models.Character{}
		models.Characters = append(models.Characters, c)

		if c.ID, err = p.getInt(d, ID); err != nil {
			return
		}

		if c.JobID, err = p.getInt(d, JobID); err != nil {
			return
		}

		if c.Name, err = p.getString(d, Name); err != nil {
			return
		}
		models.CharacterNames = append(models.CharacterNames, c.Name)

		//if pr.IsMainCharacter(c.Name) {
		models.GetParty().AddPossibleMember(&models.Member{
			CharacterID: c.ID,
			Name:        c.Name,
		})
		//}

		if c.IsEnabled, err = p.getBool(d, IsEnableCorps); err != nil {
			return
		}

		params := jo.NewOrderedMap()
		if err = p.unmarshalFrom(d, Parameter, params); err != nil {
			return
		}

		if c.Level, err = p.getInt(params, AdditionalLevel); err != nil {
			return
		}

		if c.CurrentHP, err = p.getInt(params, CurrentHP); err != nil {
			return
		}
		if c.MaxHP, err = p.getInt(params, AdditionalMaxHp); err != nil {
			return
		}

		if c.CurrentMP, err = p.getInt(params, CurrentMP); err != nil {
			return
		}
		if c.MaxMP, err = p.getInt(params, AdditionalMaxMp); err != nil {
			return
		}

		if c.Exp, err = p.getInt(d, CurrentExp); err != nil {
			return
		}

		// TODO Status

		if c.Power, err = p.getInt(params, AdditionalPower); err != nil {
			return
		}

		if c.Vitality, err = p.getInt(params, AdditionalVitality); err != nil {
			return
		}

		if c.Agility, err = p.getInt(params, AdditionalAgility); err != nil {
			return
		}

		if c.Weight, err = p.getInt(params, AdditionalWeight); err != nil {
			return
		}

		if c.Intelligence, err = p.getInt(params, AdditionalIntelligence); err != nil {
			return
		}

		if c.Spirit, err = p.getInt(params, AdditionalSpirit); err != nil {
			return
		}

		if c.Attack, err = p.getInt(params, AdditionAttack); err != nil {
			return
		}

		if c.Defense, err = p.getInt(params, AdditionalDefence); err != nil {
			return
		}

		if c.AbilityDefence, err = p.getInt(params, AdditionAbilityDefense); err != nil {
			return
		}

		if c.AbilityEvasionRate, err = p.getInt(params, AdditionAbilityEvasionRate); err != nil {
			return
		}

		if c.Magic, err = p.getInt(params, AdditionMagic); err != nil {
			return
		}

		if c.Luck, err = p.getInt(params, AdditionLuck); err != nil {
			return
		}

		if c.AccuracyRate, err = p.getInt(params, AdditionalAccuracyRate); err != nil {
			return
		}

		if c.EvasionRate, err = p.getInt(params, AdditionalEvasionRate); err != nil {
			return
		}

		if c.AbilityDisturbedRate, err = p.getInt(params, AdditionAbilityDistrurbedRate); err != nil {
			return
		}

		if c.CriticalRate, err = p.getInt(params, AdditionalCriticalRate); err != nil {
			return
		}

		if c.DamageDirmeter, err = p.getInt(params, AdditionalDamageDirmeter); err != nil {
			return
		}

		if c.AbilityDefenseRate, err = p.getInt(params, AdditionalAbilityDefenseRate); err != nil {
			return
		}

		if c.AccuracyCount, err = p.getInt(params, AdditionalAccuracyCount); err != nil {
			return
		}

		if c.EvasionCount, err = p.getInt(params, AdditionalEvasionCount); err != nil {
			return
		}

		if c.DefenseCount, err = p.getInt(params, AdditionalDefenceCount); err != nil {
			return
		}

		if c.MagicDefenseCount, err = p.getInt(params, AdditionalMagicDef); err != nil {
			return
		}
	}
	return
}

func (p *PR) loadMiscStats() (err error) {
	m := models.GetMisc()
	if m.GP, err = p.getInt(p.UserData, OwnedGil); err != nil {
		return
	}
	if m.Steps, err = p.getInt(p.UserData, Steps); err != nil {
		return
	}
	if m.EscapeCount, err = p.getInt(p.UserData, EscapeCount); err != nil {
		return
	}
	if m.BattleCount, err = p.getInt(p.UserData, BattleCount); err != nil {
		return
	}
	if m.NumberOfSaves, err = p.getInt(p.UserData, SaveCompleteCount); err != nil {
		return
	}

	if ds, ok := p.Base.GetValue(DataStorage); ok {
		jm := jo.NewOrderedMap()
		if err = jm.UnmarshalJSON([]byte(ds.(string))); err != nil {
			return
		}
		if models.GetMisc().CursedShieldFightCount, err = p.getIntFromSlice(jm, "global"); err != nil {
			return
		}
	}

	if m.OpenedChestCount, err = p.getInt(p.UserData, OpenChestCount); err != nil {
		return
	}
	if m.IsCompleteFlag, err = p.getFlag(p.Base, IsCompleteFlag); err != nil {
		return
	}
	if m.PlayTime, err = p.getFloat(p.UserData, PlayTime); err != nil {
		return
	}

	if global.GetSaveType() == global.One {
		var sl interface{}
		if sl, err = p.getFromTarget(p.UserData, OwnendCrystalFlags); err != nil {
			return
		}
		for i, j := range sl.([]interface{}) {
			m.OwnedCrystals[i] = j.(bool)
		}
	}

	return
}

func (p *PR) loadInventory(key string, inventory *models.Inventory) (err error) {
	var (
		sl  interface{}
		row models.Row
	)
	if sl, err = p.getFromTarget(p.UserData, key); err != nil {
		return
	}
	inventory.Reset()
	for i, r := range sl.([]interface{}) {
		if err = json.Unmarshal([]byte(r.(string)), &row); err != nil {
			return
		}
		inventory.Set(i, row)
	}
	return nil
}

func (p *PR) loadMapData() (err error) {
	md := models.GetMapData()
	if md.MapID, err = p.getInt(p.MapData, MapID); err != nil {
		return
	}
	if md.PointIn, err = p.getInt(p.MapData, PointIn); err != nil {
		return
	}
	if md.TransportationID, err = p.getInt(p.MapData, TransportationID); err != nil {
		return
	}
	if md.CarryingHoverShip, err = p.getBool(p.MapData, CarryingHoverShip); err != nil {
		return
	}
	if md.PlayableCharacterCorpsID, err = p.getInt(p.MapData, PlayableCharacterCorpsId); err != nil {
		return
	}

	pe := jo.NewOrderedMap()
	if err = p.unmarshalFrom(p.MapData, PlayerEntity, pe); err != nil {
		return
	}
	var pos *jo.OrderedMap
	if pos = pe.Get(PlayerPosition).(*jo.OrderedMap); pos == nil {
		err = errors.New("unable to get transportation position")
		return
	}
	if md.Player.X, err = p.getFloat(pos, "x"); err != nil {
		return
	}
	if md.Player.Y, err = p.getFloat(pos, "y"); err != nil {
		return
	}
	if md.Player.Z, err = p.getFloat(pos, "z"); err != nil {
		return
	}
	if md.PlayerDirection, err = p.getInt(pe, PlayerDirection); err != nil {
		return
	}

	gps := jo.NewOrderedMap()
	if err = p.unmarshalFrom(p.MapData, GpsData, gps); err != nil {
		return
	}
	if md.Gps.MapID, err = p.getInt(gps, GpsDataMapID); err != nil {
		return
	}
	if md.Gps.AreaID, err = p.getInt(gps, GpsDataAreaID); err != nil {
		return
	}
	if md.Gps.GpsID, err = p.getInt(gps, GpsDataID); err != nil {
		return
	}
	if md.Gps.Width, err = p.getInt(gps, GpsDataWidth); err != nil {
		return
	}
	if md.Gps.Height, err = p.getInt(gps, GpsDataHeight); err != nil {
		return
	}
	return
}

func (p *PR) loadTransportation() (err error) {
	var sl interface{}
	if sl, err = p.getFromTarget(p.UserData, OwnedTransportationList); err != nil {
		return
	}
	models.Transportations = make([]*models.Transportation, len(sl.([]interface{})))
	for index, i := range sl.([]interface{}) {
		om := jo.NewOrderedMap()
		if err = om.UnmarshalJSON([]byte(i.(string))); err != nil {
			return
		}
		t := &models.Transportation{}
		if t.ID, err = p.getInt(om, TransID); err != nil {
			return
		}
		if t.MapID, err = p.getInt(om, TransMapID); err != nil {
			return
		}
		if t.Direction, err = p.getInt(om, TransDirection); err != nil {
			return
		}
		if t.TimeStampTicks, err = p.getUint(om, TransTimeStampTicks); err != nil {
			return
		}

		var pos *jo.OrderedMap
		if pos = om.Get(TransPosition).(*jo.OrderedMap); pos == nil {
			err = errors.New("unable to get transportation position")
			return
		}
		if t.Position.X, err = p.getFloat(pos, "x"); err != nil {
			return
		}
		if t.Position.Y, err = p.getFloat(pos, "y"); err != nil {
			return
		}
		if t.Position.Z, err = p.getFloat(pos, "z"); err != nil {
			return
		}

		t.Enabled = t.TimeStampTicks > 0 && t.MapID > 0 && t.Position.X > 0 && t.Position.Y > 0 && t.Position.Z > 0

		models.Transportations[index] = t
	}
	return
}

func (p *PR) getString(c *jo.OrderedMap, key string) (s string, err error) {
	j, ok := c.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}
	if s, ok = j.(string); !ok {
		err = fmt.Errorf("unable to parse field %s value %v ", key, j)
	}
	return
}

func (p *PR) getBool(c *jo.OrderedMap, key string) (b bool, err error) {
	j, ok := c.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}
	if b, ok = j.(bool); !ok {
		err = fmt.Errorf("unable to parse field %s value %v", key, j)
	}
	return
}

func (p *PR) getInt(c *jo.OrderedMap, key string) (i int, err error) {
	j, ok := c.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}

	k := reflect.ValueOf(j)
	switch k.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(k.Int()), nil
	case reflect.Float32, reflect.Float64:
		return int(k.Float()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int(k.Uint()), nil
	case reflect.String:
		var l int64
		l, err = strconv.ParseInt(k.String(), 10, 32)
		if err == nil {
			i = int(l)
		}
	default:
		err = fmt.Errorf("unable to parse field %s value %v ", key, j)
	}
	return
}

func (p *PR) getUint(c *jo.OrderedMap, key string) (i uint64, err error) {
	j, ok := c.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}

	k := reflect.ValueOf(j)
	switch k.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(k.Int()), nil
	case reflect.Float32, reflect.Float64:
		return uint64(k.Float()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uint64(k.Uint()), nil
	case reflect.String:
		var l int64
		l, err = strconv.ParseInt(k.String(), 10, 64)
		if err == nil {
			i = uint64(l)
		}
	default:
		err = fmt.Errorf("unable to parse field %s value %v ", key, j)
	}
	return
}

func (p *PR) getFloat(c *jo.OrderedMap, key string) (f float64, err error) {
	j, ok := c.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}

	k := reflect.ValueOf(j)
	switch k.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return k.Float(), nil
	case reflect.Float32, reflect.Float64:
		return k.Float(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return k.Float(), nil
	case reflect.String:
		f, err = strconv.ParseFloat(k.String(), 64)
	default:
		err = fmt.Errorf("unable to parse field %s value %v ", key, j)
	}
	return
}

func (p *PR) getJsonInts(data *jo.OrderedMap, key string) (ints []interface{}, err error) {
	j, ok := data.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}
	ints, ok = j.([]interface{})
	if !ok {
		err = fmt.Errorf("unable to load %s", key)
	}
	return
}

func (p *PR) getFlag(m *jo.OrderedMap, key string) (b bool, err error) {
	var i int
	if i, err = p.getInt(m, key); err != nil {
		return
	}
	if i == 0 {
		b = false
	} else {
		b = true
	}
	return
}

func (p *PR) getIntFromSlice(from *jo.OrderedMap, key string) (v int, err error) {
	var (
		i, ok = from.GetValue(key)
		sl    []interface{}
		i64   int64
	)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
		return
	}

	if sl, ok = i.([]interface{}); !ok || len(sl) < 9 {
		err = fmt.Errorf("unable to load cursed shield battle count")
		return
	}

	if i64, err = sl[9].(json.Number).Int64(); err != nil {
		return
	}
	v = int(i64)
	return
}

func (p *PR) unmarshalFrom(from *jo.OrderedMap, key string, m *jo.OrderedMap) (err error) {
	i, ok := from.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}
	switch t := i.(type) {
	case string:
		err = m.UnmarshalJSON([]byte(t))
	default:
		err = fmt.Errorf("cannot unmarshal unkown type %v", i)
	}
	return
}

func (p *PR) unmarshal(i interface{}, m *map[string]interface{}) error {
	//s := strings.ReplaceAll(i.(string), `\\"`, `\"`)
	return json.Unmarshal([]byte(i.(string)), m)
}

func (p *PR) unmarshalEquipment(m *jo.OrderedMap) (idCounts []idCount, err error) {
	i, ok := m.GetValue(EquipmentList)
	if !ok {
		return nil, fmt.Errorf("%s not found", EquipmentList)
	}

	eq := jo.NewOrderedMap()
	if err = eq.UnmarshalJSON([]byte(i.(string))); err != nil {
		return
	}

	if i, ok = eq.GetValue("values"); ok && i != nil {
		idCounts = make([]idCount, len(i.([]interface{})))
		for j, v := range i.([]interface{}) {
			if err = json.Unmarshal([]byte(v.(string)), &idCounts[j]); err != nil {
				return
			}
		}
	}
	return
}

func (p *PR) getFromTarget(data *jo.OrderedMap, key string) (i interface{}, err error) {
	var (
		slTarget = jo.NewOrderedMap()
		ok       bool
	)
	if err = p.unmarshalFrom(data, key, slTarget); err != nil {
		return
	}
	if i, ok = slTarget.GetValue(targetKey); !ok {
		err = fmt.Errorf("unable to find %s", targetKey)
	}
	return
}

func (p *PR) fixEscapeCharsForLoad(s string) string {
	var (
		sb    strings.Builder
		going = false
		count int
		found = make(map[int]string)
	)
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' {
			if !going {
				going = true
			}
			sb.WriteByte('\\')
			count++
		} else {
			if going {
				going = false
				sb.WriteByte('"')
				found[count] = sb.String()
				count = 0
				sb.Reset()
			}
		}
	}

	sorted := make([]int, 0, len(found))
	for k, _ := range found {
		sorted = append(sorted, k)
	}
	sort.Ints(sorted)

	for i := len(sorted) - 1; i >= 0; i-- {
		count = sorted[i]
		v, _ := found[count]
		sb.Reset()
		for j := 0; j < count; j++ {
			sb.WriteByte('~')
		}
		sb.WriteByte('"')
		found[count] = sb.String()
		s = strings.ReplaceAll(s, v, sb.String())
	}

	for i := len(sorted) - 1; i >= 0; i-- {
		count = sorted[i]
		v, _ := found[count]
		sb.Reset()
		half := (len(v) - 1) / 2
		for j := 0; j < half; j++ {
			sb.WriteByte('\\')
		}
		if sb.Len() == 0 {
			sb.WriteByte('\\')
		}
		sb.WriteByte('"')
		s = strings.ReplaceAll(s, v, sb.String())
	}
	return s
}

// {"keys":[1,2,3,4,5,6],"values":["{\\"contentId\\":113,\\"count\\":2}","{\\"contentId\\":214,\\"count\\":1}","{\\"contentId\\":268,\\"count\\":1}","{\\"contentId\\":233,\\"count\\":8}","{\\"contentId\\":200,\\"count\\":67}","{\\"contentId\\":315,\\"count\\":1}"]}
type equipment struct {
	Keys   []int    `json:"keys"`
	Values []string `json:"values"`
}

type idCount struct {
	ContentID int `json:"contentId"`
	Count     int `json:"count"`
}

func (p *PR) execLoad(fileName string, omitFirstBytes bool) ([]byte, error) {
	if _, err := os.Stat("pr_io"); err != nil {
		if err = p.downloadPyExe(); err != nil {
			return nil, err
		}
	}
	if _, err := os.Stat("pr_io.zip"); err != nil {
		_ = os.Remove("pr_io.zip")
	}

	s := "0"
	if omitFirstBytes {
		s = "1"
	}

	path := strings.ReplaceAll(filepath.Join(global.PWD, "pr_io"), "\\", "/")
	cmd := exec.Command("cmd", "/C", "pr_io.exe", "deobfuscateFile", fileName, s)
	cmd.Dir = path
	return cmd.Output()
}

func handleCmdError(out []byte, err error) error {
	if e, ok := err.(*exec.ExitError); ok {
		return fmt.Errorf("failed to load file: " + string(e.Stderr))
	}
	return fmt.Errorf("failed to load file:\n%s\n%s", err.Error(), string(out))
}

func (p *PR) loadBase(s string) (err error) {
	return p.Base.UnmarshalJSON([]byte(s))
}

func (p *PR) getSaveData(s string) (string, error) {
	var (
		start int
		end   = strings.Index(s, `,"clearFlag`)
	)
	if end == -1 {
		return "", errors.New("unable to load file. Please try resaving to a new unused game slot and try loading that slot instead")
	}
	for start < len(s) && s[start] != '{' {
		start++
	}
	end = len(s) - 1
	for end >= 0 && s[end] != '}' {
		end--
	}
	if end+1 >= len(s) {
		return "", errors.New("unable to load file. Please try resaving to a new unused game slot and try loading that slot instead")
	}
	return s[start : end+1], nil // + `,"playTime":0.0,"clearFlag":0}`, nil
}

func (p *PR) readFile(fileName string) (out []byte, err error) {
	if out, err = p.execLoad(fileName, true); err != nil {
		e1 := handleCmdError(out, err)
		if out, err = p.execLoad(fileName, false); err != nil {
			err = e1
			return
		}
	}
	return
}

func (p *PR) uncheckAll(rages []*models.NameValueChecked) {
	for _, r := range rages {
		r.Checked = false
	}
}

func (p *PR) replaceUnicodeNames(b []byte, names *[]unicodeNameReplace) ([]byte, error) {
	for i := 0; i < len(b)-10; i++ {
		if b[i] == '"' && b[i+1] == 'n' && b[i+2] == 'a' && b[i+3] == 'm' && b[i+4] == 'e' && b[i+5] == '\\' {
			i += 5
			for i < len(b)-1 && (b[i] == '\\' || b[i] == ':' || b[i] == ' ' || b[i] == '"') {
				i++
			}
			if b[i-1] != '"' {
				i--
			}

			// Start of name is i

			isUnicode := false
			for j := i; j < len(b)-50 && b[j] != '"' && !(b[j] == '\\' && b[j+1] == '\\'); j++ {
				if b[j] == '\\' && b[j+1] == 'x' {
					isUnicode = true
				}
			}
			if !isUnicode {
				continue
			}

			var original, replaced []byte
			for i < len(b)-1 && b[i] != '"' && !(b[i] == '\\' && b[i+1] == '\\') {
				original = append(original, b[i])
				replaced = append(replaced, '~')
				b[i] = '~'
				i++
			}
			r := unicodeNameReplace{
				Replaced: string(replaced),
			}
			var err error
			sb := strings.Builder{}
			for j := 0; j < len(original)-1; j++ {
				var char []byte
				if original[j] == '\\' && original[j+1] == 'x' {
					char = append(char, original[j])
					for j++; j < len(original); j++ {
						if original[j] == '\\' {
							break
						}
						char = append(char, original[j])
					}
					j--
					var s string
					if s, err = strconv.Unquote(strings.Replace(strconv.Quote(string(char)), `\\x`, `\x`, -1)); err != nil {
						return nil, err
					}
					sb.WriteString(s)
				} else {
					sb.WriteString(string(original[j]))
				}
			}
			r.Original = sb.String()
			*names = append(*names, r)
		}
	}
	return b, nil
}

type unicodeNameReplace struct {
	Original string
	Replaced string
}

func (p *PR) downloadPyExe() error {
	var (
		resp, err = http.Get("https://github.com/KiameV/pr_save_io/releases/download/latest/pr_io.zip")
		out       *os.File
		r         *zip.ReadCloser
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("status", resp.Status)
	if resp.StatusCode != 200 {
		return errors.New("failed to download the save file reader")
	}

	// Create the file
	if out, err = os.Create("pr_io.zip"); err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	if _, err = io.Copy(out, resp.Body); err != nil {
		return err
	}

	if r, err = zip.OpenReader("pr_io.zip"); err != nil {
		return err
	}
	defer func() { _ = r.Close() }()

	if err = os.Mkdir("./pr_io", 0755); err != nil {
		return err
	}

	// Closure to address file descriptors issue with all the deferred .Close() methods
	for _, f := range r.File {
		if err = extractArchiveFile(".", f); err != nil {
			return err
		}
	}
	return nil
}

func extractArchiveFile(dest string, f *zip.File) (err error) {
	var (
		rc   io.ReadCloser
		file *os.File
		path string
	)
	if rc, err = f.Open(); err != nil {
		return
	}
	defer func() { _ = rc.Close() }()

	path = filepath.Join(dest, f.Name)
	// Check for ZipSlip (Directory traversal)
	path = strings.ReplaceAll(path, "..", "")

	if f.FileInfo().IsDir() {
		if err = os.MkdirAll(path, f.Mode()); err != nil {
			return
		}
	} else {
		if err = os.MkdirAll(filepath.Dir(path), f.Mode()); err != nil {
			return
		}
		if file, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode()); err != nil {
			return
		}
		defer func() { _ = file.Close() }()
		if _, err = io.Copy(file, rc); err != nil {
			return
		}
	}
	return
}

/*
d := []byte(s)
		for i, c := range d {
			if c == 'x' && d[i-1] == '\\' {
				d[i] = '^'
				d[i-1] = '^'
			}
		}

func (p *PR) fixFile(s string) (bool, string) {
	if i := strings.Index(s, "clearFlag"); i != -1 {
		c := s[i+9]
		if c != ':' {
			cc := s[i+10]
			if cc >= 48 && c <= 57 {
				cc -= 48
			} else {
				cc = 0
			}
			s = s[:i+9] + fmt.Sprintf(`":%d}`, cc)
		}
		return true, s
	} else if i = strings.Index(s, `"playTime`); i != -1 && s[i+4] >= 48 && s[i+4] <= 57 {
		s = s[0:i] + `"playTime":0.0,"clearFlag":0}`
		return true, s
	}
	return false, s
}*/
