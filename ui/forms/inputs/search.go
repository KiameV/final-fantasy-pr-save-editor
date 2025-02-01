package inputs

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/models"
	"pixel-remastered-save-editor/models/core/costs/ff1"
	"pixel-remastered-save-editor/models/core/costs/ff2"
	"pixel-remastered-save-editor/models/core/costs/ff3"
	"pixel-remastered-save-editor/models/core/costs/ff4"
	"pixel-remastered-save-editor/models/core/costs/ff5"
	"pixel-remastered-save-editor/models/core/costs/ff6"
	"pixel-remastered-save-editor/models/finder"
)

type (
	Searches struct {
		Abilities      *Search
		Characters     *Search
		Commands       *Search
		Equipment      *Search
		Items          *Search
		ImportantItems *Search
		ItemsEquipment *Search
		Jobs           *Search
		Maps           *Search
		Bestiary       *Search
	}
	Search struct {
		Entries *[]models.NameValue
		Text    *widget.TextGrid
		Input   *widget.Entry
		Results *widget.TextGrid
	}
)

var _s *Searches

func GetSearches() *Searches {
	return _s
}

func Load(game global.Game) {
	_s = &Searches{
		Characters: NewSearch(finder.AllCharacters()),
	}
	if game == global.One {
		_s.Abilities = NewSearch(ff1.Abilities)
		_s.Commands = NewSearch(ff1.Commands)
		_s.Equipment = NewSearch(ff1.Weapons, ff1.Shields, ff1.Armors, ff1.Helmets, ff1.Gloves)
		_s.Items = NewSearch(ff1.Items)
		_s.ImportantItems = NewSearch(ff1.ImportantItems)
		_s.ItemsEquipment = NewSearch(ff1.Items, ff1.Weapons, ff1.Shields, ff1.Armors, ff1.Helmets, ff1.Gloves)
		_s.Jobs = NewSearch(ff1.Jobs)
		_s.Maps = NewSearch(ff1.Maps)
		_s.Bestiary = NewSearch(ff1.Bestiary)
	} else if game == global.Two {
		_s.Abilities = newSearchFF2(ff2.Abilities)
		_s.Commands = NewSearch(ff2.Commands)
		_s.Equipment = NewSearch(ff2.Weapons, ff2.Shields, ff2.Armors, ff2.Helmets, ff2.Gloves)
		_s.ItemsEquipment = NewSearch(ff2.Items, ff2.Weapons, ff2.Shields, ff2.Armors, ff2.Helmets, ff2.Gloves)
		_s.Items = NewSearch(ff2.Items)
		_s.Jobs = NewSearch(ff2.Jobs)
		_s.Maps = NewSearch(ff2.Maps)
		_s.Bestiary = NewSearch(ff2.Bestiary)
	} else if game == global.Three {
		_s.Abilities = NewSearch(ff3.Abilities, ff3.WhiteMagic, ff3.BlackMagic, ff3.SummonMagic)
		_s.Commands = NewSearch(ff3.Commands)
		_s.Equipment = NewSearch(ff3.Weapons, ff3.Shields, ff3.Armors, ff3.Helmets, ff3.Hands)
		_s.Items = NewSearch(ff3.Items)
		_s.ImportantItems = NewSearch(ff3.ImportantItems)
		_s.ItemsEquipment = NewSearch(ff3.Items, ff3.Weapons, ff3.Shields, ff3.Armors, ff3.Helmets, ff3.Hands)
		_s.Jobs = NewSearch(ff3.Jobs)
		_s.Maps = NewSearch(ff3.Maps)
		_s.Bestiary = NewSearch(ff3.Bestiary)
	} else if game == global.Four {
		_s.Abilities = NewSearch(ff4.Abilities, ff4.WhiteMagic, ff4.BlackMagic, ff4.SummonMagic)
		_s.Commands = NewSearch(ff4.Commands)
		_s.Equipment = NewSearch(ff4.Weapons, ff4.Shields, ff4.Armors, ff4.Helmets, ff4.Hands)
		_s.Items = NewSearch(ff4.Items)
		_s.ImportantItems = NewSearch(ff4.ImportantItems)
		_s.ItemsEquipment = NewSearch(ff4.Items, ff4.Weapons, ff4.Shields, ff4.Armors, ff4.Helmets, ff4.Hands)
		_s.Jobs = NewSearch(ff4.Jobs)
		_s.Maps = NewSearch(ff4.Maps)
		_s.Bestiary = NewSearch(ff4.Bestiary)
	} else if game == global.Five {
		_s.Abilities = NewSearch(ff5.Abilities, ff5.Spellblade, ff5.WhiteMagic, ff5.BlackMagic, ff5.SummonMagic, ff5.TimeMagic, ff5.BlueMagic, ff5.Songs)
		_s.Commands = NewSearch(ff5.Commands)
		_s.Equipment = NewSearch(ff5.Weapons, ff5.Shields, ff5.Armors, ff5.Helmets, ff5.Hands)
		_s.Items = NewSearch(ff5.Items)
		_s.ItemsEquipment = NewSearch(ff5.Items, ff5.Weapons, ff5.Shields, ff5.Armors, ff5.Helmets, ff5.Hands)
		_s.Jobs = NewSearch(ff5.Jobs)
		_s.Maps = NewSearch(ff5.Maps)
	} else { // Six
		_s.Abilities = NewSearch(ff6.Blitzes, ff6.Bushidos, ff6.Dances, ff6.Lores, ff6.Rages, ff6.Magic)
		_s.Commands = NewSearch(ff6.Commands)
		_s.Equipment = NewSearch(ff6.Weapons, ff6.Shields, ff6.Armors, ff6.Helmets, ff6.Hands)
		_s.Items = NewSearch(ff6.Items)
		_s.ItemsEquipment = NewSearch(ff6.Items, ff6.Weapons, ff6.Shields, ff6.Armors, ff6.Helmets, ff6.Hands)
		_s.Jobs = NewSearch(ff6.Jobs)
		_s.Maps = NewSearch(ff6.Maps)
		_s.Bestiary = NewSearch(ff6.Bestiary)
	}
}

func NewSearch(nvss ...[]models.NameValue) *Search {
	s := &Search{
		Text:    widget.NewTextGrid(),
		Input:   widget.NewEntry(),
		Results: widget.NewTextGrid(),
		Entries: &nvss[0],
	}
	s.onFilterChanged(nvss...)
	var sb strings.Builder
	for _, nvs := range nvss {
		for _, nv := range nvs {
			sb.WriteString(fmt.Sprintf("%d: %s\n", nv.Value, nv.Name))
		}
	}
	s.Text.SetText(sb.String())
	return s
}

func newSearchFF2(nvss ...[]models.NameValue) *Search {
	s := &Search{
		Text:    widget.NewTextGrid(),
		Input:   widget.NewEntry(),
		Results: widget.NewTextGrid(),
	}
	s.onFilterChanged(nvss...)
	m := make(map[string]bool)
	var r []struct {
		name     string
		from, to int
	}
	for _, nvs := range nvss {
		for _, nv := range nvs {
			n := strings.Split(nv.Name, " ")[0]
			if _, ok := m[n]; !ok {
				m[n] = true
				r = append(r, struct {
					name     string
					from, to int
				}{name: n, from: nv.Value, to: nv.Value + 15})
			}
		}
	}
	var sb strings.Builder
	for _, i := range r {
		sb.WriteString(fmt.Sprintf("%d-%d: %s\n", i.from, i.to, i.name))
	}
	s.Text.SetText(sb.String())
	return s
}

func (s *Search) onFilterChanged(nvss ...[]models.NameValue) {
	s.Input.OnChanged = func(text string) {
		if len(text) >= 2 {
			text = strings.ToLower(text)
			var sb strings.Builder
			for _, nvs := range nvss {
				for _, nv := range nvs {
					if strings.Contains(strings.ToLower(nv.Name), text) {
						sb.WriteString(fmt.Sprintf("%d: %s\n", nv.Value, nv.Name))
					}
				}
			}
			s.Results.SetText(sb.String())
		} else {
			s.Results.SetText("")
		}
	}
}

func (s *Search) Filter() fyne.CanvasObject {
	if s == nil {
		return container.NewStack()
	}
	return container.NewBorder(s.Input, nil, nil, nil, container.NewVScroll(s.Results))
}

func (s *Search) Fields() fyne.CanvasObject {
	if s == nil {
		return container.NewStack()
	}
	return container.NewVScroll(s.Text)
}

func (s *Search) Count() (i int) {
	if s.Entries != nil {
		i = len(*s.Entries)
	}
	return
}
