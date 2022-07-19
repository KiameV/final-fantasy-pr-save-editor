package models

import (
	"sort"
)

var EmptyPartyMember = &Member{
	CharacterID: 0,
	Name:        "[Empty]",
}

type Member struct {
	CharacterID int `json:"characterId"`
	Name        string
	//EnableEquipment bool
}

var party *Party

type Party struct {
	Members       [4]*Member
	Possible      map[string]*Member
	PossibleNames []string
	//PossibleNamesWithNPCs []string
	//Enabled bool
	//IncludeNPCs bool
}

func GetParty() *Party {
	if party == nil {
		party = &Party{}
		party.Clear()
	}
	return party
}

func (p *Party) Clear() {
	p.Possible = make(map[string]*Member)
	p.PossibleNames = make([]string, 0, 40)
	//p.PossibleNamesWithNPCs = make([]string, 0, 40)
	p.AddPossibleMember(EmptyPartyMember)
}

func (p *Party) AddPossibleMember(m *Member) {
	c, found := GetCharacter(m.CharacterID)
	if !found {
		return
	}
	if _, found = p.Possible[c.Name]; !found || c.IsEnabled {
		p.Possible[c.Name] = m
		if !found {
			p.PossibleNames = append(p.PossibleNames, m.Name)
			sort.Strings(p.PossibleNames)
		}
	}
}

func (p *Party) GetPossibleIndex(m *Member) int {
	names := p.PossibleNames
	//if p.IncludeNPCs {
	//	names = p.PossibleNamesWithNPCs
	//}

	for i, po := range names {
		if m.Name == po {
			return i
		}
	}
	return 0
}

func (p *Party) SetMember(slot int, member *Member) {
	p.Members[slot] = member
}
