package finder

import (
	"cmp"
	"slices"

	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/models"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/core/costs/ff1"
	"pixel-remastered-save-editor/models/core/costs/ff2"
	"pixel-remastered-save-editor/models/core/costs/ff3"
	"pixel-remastered-save-editor/models/core/costs/ff4"
	"pixel-remastered-save-editor/models/core/costs/ff5"
	"pixel-remastered-save-editor/models/core/costs/ff6"
)

type (
	Find    func(int) (string, bool)
	Finders interface {
		Abilities(int) (string, bool)
		Characters(int) (string, bool)
		Commands(int) (string, bool)
		Items(int) (string, bool)
		ImportantItems(int) (string, bool)
		Jobs(int) (string, bool)
		Maps(int) (string, bool)
	}
	finders struct {
		abilities  map[int]string
		characters map[int]string
		commands   map[int]string
		items      map[int]string
		important  map[int]string
		jobs       map[int]string
		maps       map[int]string
		bestiary   map[int]string
	}
)

var (
	singletonFinder *finders
)

func Load(game global.Game, characters []*core.Character) {
	if game == global.One {
		singletonFinder = &finders{
			abilities: nameLookup(ff1.Abilities),
			commands:  nameLookup(ff1.Commands),
			important: nameLookup(ff1.ImportantItems),
			items:     nameLookup(ff1.Items, ff1.Weapons, ff1.Shields, ff1.Armors, ff1.Helmets, ff1.Gloves),
			jobs:      nameLookup(ff1.Jobs),
			maps:      nameLookup(),
			bestiary:  nameLookup(ff1.Bestiary),
		}
	} else if game == global.Two {
		singletonFinder = &finders{
			abilities: nameLookup(ff2.Abilities),
			commands:  nameLookup(ff2.Commands),
			important: nameLookup(),
			items:     nameLookup(ff2.Items, ff2.Weapons, ff2.Shields, ff2.Armors, ff2.Helmets, ff2.Gloves),
			jobs:      nameLookup(ff2.Jobs),
			maps:      nameLookup(ff2.Maps),
			bestiary:  nameLookup(ff2.Bestiary),
		}
	} else if game == global.Three {
		singletonFinder = &finders{
			abilities: nameLookup(ff3.Abilities, ff3.WhiteMagic, ff3.BlackMagic, ff3.SummonMagic),
			commands:  nameLookup(ff3.Commands),
			important: nameLookup(ff3.ImportantItems),
			items:     nameLookup(ff3.Items, ff3.Weapons, ff3.Shields, ff3.Armors, ff3.Helmets, ff3.Hands),
			jobs:      nameLookup(ff3.Jobs),
			maps:      nameLookup(ff3.Maps),
			bestiary:  nameLookup(ff3.Bestiary),
		}
	} else if game == global.Four {
		singletonFinder = &finders{
			abilities: nameLookup(ff4.Abilities, ff4.WhiteMagic, ff4.BlackMagic, ff4.SummonMagic),
			commands:  nameLookup(ff4.Commands),
			important: nameLookup(ff4.ImportantItems),
			items:     nameLookup(ff4.Items, ff4.Weapons, ff4.Shields, ff4.Armors, ff4.Helmets, ff4.Hands),
			jobs:      nameLookup(ff4.Jobs),
			maps:      nameLookup(ff4.Maps),
			bestiary:  nameLookup(ff4.Bestiary),
		}
	} else if game == global.Five {
		singletonFinder = &finders{
			abilities: nameLookup(ff5.Abilities, ff5.Spellblade, ff5.WhiteMagic, ff5.BlackMagic, ff5.SummonMagic, ff5.TimeMagic, ff5.BlueMagic, ff5.Songs),
			commands:  nameLookup(ff5.Commands),
			important: nameLookup(ff5.ImportantItems),
			items:     nameLookup(ff5.Items, ff5.Weapons, ff5.Shields, ff5.Armors, ff5.Helmets, ff5.Hands),
			jobs:      nameLookup(ff5.Jobs),
			maps:      nameLookup(ff5.Maps),
			bestiary:  nameLookup(ff5.Bestiary),
		}
	} else { // Six
		singletonFinder = &finders{
			abilities: nameLookup(ff6.Blitzes, ff6.Dances, ff6.Lores, ff6.Bushidos, ff6.Rages, ff6.Magic),
			commands:  nameLookup(ff6.Commands),
			important: nameLookup(ff6.ImportantItems),
			items:     nameLookup(ff6.Items, ff6.Weapons, ff6.Shields, ff6.Armors, ff6.Helmets, ff6.Hands),
			jobs:      nameLookup(ff6.Jobs),
			maps:      nameLookup(ff6.Maps),
			bestiary:  nameLookup(ff6.Bestiary),
		}
	}
	singletonFinder.characters = make(map[int]string)
	for _, c := range characters {
		singletonFinder.characters[c.Base.ID] = c.Base.Name
	}
}

func New(abilities []models.NameValue) Find {
	l := nameLookup(abilities)
	return func(i int) (s string, b bool) {
		s, b = l[i]
		return
	}
}

func Abilities(i int) (s string, b bool) {
	return singletonFinder.Abilities(i)
}

func (f finders) Abilities(i int) (s string, b bool) {
	s, b = f.abilities[i]
	return
}

func Characters(i int) (string, bool) {
	return singletonFinder.Characters(i)
}

func (f finders) Characters(i int) (s string, b bool) {
	s, b = f.characters[i]
	return
}

func Commands(i int) (s string, b bool) {
	return singletonFinder.Commands(i)
}

func (f finders) Commands(i int) (s string, b bool) {
	s, b = f.commands[i]
	return
}

func Items(i int) (s string, b bool) {
	return singletonFinder.Items(i)
}

func (f finders) Items(i int) (s string, b bool) {
	s, b = f.items[i]
	return
}

func ImportantItems(i int) (s string, b bool) {
	return singletonFinder.ImportantItems(i)
}

func (f finders) ImportantItems(i int) (s string, b bool) {
	s, b = f.important[i]
	return
}

func Jobs(i int) (s string, b bool) {
	return singletonFinder.Jobs(i)
}

func (f finders) Jobs(i int) (s string, b bool) {
	s, b = f.jobs[i]
	return
}

func Maps(i int) (s string, b bool) {
	return singletonFinder.Maps(i)
}

func (f finders) Maps(i int) (s string, b bool) {
	s, b = f.maps[i]
	return
}

func Bestiary(i int) (s string, b bool) { return singletonFinder.Bestiary(i) }

func (f finders) Bestiary(i int) (s string, b bool) {
	s, b = f.bestiary[i]
	return
}

func Get() Finders {
	return singletonFinder
}

func nameLookup(args ...[]models.NameValue) map[int]string {
	m := make(map[int]string)
	for _, a := range args {
		for _, i := range a {
			m[i.Value] = i.Name
		}
	}
	return m
}

func AllCharacters() (v []models.NameValue) {
	v = make([]models.NameValue, 0, len(singletonFinder.characters))
	for i, n := range singletonFinder.characters {
		v = append(v, models.NewNameValue(n, i))
	}
	slices.SortFunc(v, func(i, j models.NameValue) int {
		return cmp.Compare(i.Value, j.Value)
	})
	return
}

func AllJobs() (v map[int]string) {
	return singletonFinder.jobs
}
