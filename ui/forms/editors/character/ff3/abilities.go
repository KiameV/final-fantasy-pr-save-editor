package ff3

import (
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/ui/forms/editors/character/ff1"
)

type Abilities ff1.Abilities

func NewAbilities(c *core.Character) *Abilities {
	a := ff1.NewAbilities(c)
	return (*Abilities)(a)
}
