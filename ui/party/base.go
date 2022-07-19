package party

import (
	"fmt"
	"github.com/aarzilli/nucular"
	"pr_save_editor/models"
	"pr_save_editor/ui"
)

type partyUI struct {
	selected [4]int
}

func NewUI() ui.UI {
	return &partyUI{}
}

func (u *partyUI) Draw(w *nucular.Window) {
	w.Row(10).Static()
	for i := 0; i < 4; i++ {
		u.drawRow(w, i, &u.selected[i])
	}
}

func (u *partyUI) drawRow(w *nucular.Window, slot int, selected *int) {
	p := models.GetParty()
	w.Row(24).Static(60, 200, 200, 60)
	w.Label(fmt.Sprintf("Member: %d", slot+1), "LC")
	if i := w.ComboSimple(p.PossibleNames, *selected, 12); i != *selected {
		*selected = i
		p.Members[slot] = p.Possible[p.PossibleNames[i]]
	}
	if *selected > 0 {
		pn := p.PossibleNames[*selected]
		if pn != models.EmptyPartyMember.Name {
			c, found := models.GetCharacterByName(pn)
			if !found {
				// TODO
			}
			if !c.IsEnabled {
				w.CheckboxText("Enable?", &c.IsEnabled)
			}
		}
	}
}

func (u *partyUI) Refresh() {
	p := models.GetParty()
	for i, m := range p.Members {
		if m != nil {
			u.selected[i] = p.GetPossibleIndex(m)
		}
	}
}

func (u *partyUI) Name() string {
	return "Party"
}

func (u *partyUI) Behavior() ui.Behavior {
	return ui.Show
}
