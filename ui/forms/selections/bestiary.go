package selections

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/finder"
	"pixel-remastered-save-editor/ui/forms/editors/bestiary"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Bestiary struct {
		widget.BaseWidget
		bestiary *core.Bestiary
	}
)

func NewBestiary(bestiary *core.Bestiary) (b *Bestiary) {
	b = &Bestiary{
		bestiary: bestiary,
	}
	b.ExtendBaseWidget(b)
	return
}

func (b *Bestiary) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(
		bestiary.NewCore(b.bestiary, finder.Bestiary, inputs.GetSearches().Bestiary),
	)

}
