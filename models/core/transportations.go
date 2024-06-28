package core

import (
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/save"
)

type (
	Transportations struct {
		Transportations []*save.OwnedTransportation
	}
)

func NewTransportations(game global.Game, ud *save.UserData) (t *Transportations, err error) {
	t = &Transportations{}
	t.Transportations, err = ud.OwnedTransportationList()
	return
}

func (t *Transportations) ToSave(ud *save.UserData) error {
	return ud.SetOwnedTransportationList(t.Transportations)
}
