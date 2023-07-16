package widgets

import (
	"math"

	"github.com/aarzilli/nucular"
	"pr_save_editor/consts"
)

var (
	name       nucular.TextEditor
	nameResult []string
)

func init() {
	name.Flags = nucular.EditField
	name.Maxlen = 8
	name.SingleLine = true
}

func AddItemFinder(w *nucular.Window, itemSet ...[]*consts.Item) {
	if sw := w.GroupBegin("Item Finder", 0); sw != nil {
		sw.Row(24).Static(100, 80)
		sw.Label("Find By Name:", "LC")
		if e := name.Edit(sw); e == nucular.EditActive || e == nucular.EditCommitted {
			l := len(name.Buffer)
			if l == 0 || l >= 2 {
				nameResult = nameResult[:0]
			}
			if l >= 2 {
				nameResult = consts.Search(string(name.Buffer), itemSet...)
			}
		}

		for _, s := range nameResult {
			sw.Row(20).Static(150)
			sw.Label(s, "LC")
		}
		sw.GroupEnd()
	}
}

func GetTime(input int) (hours int, minutes int) {
	hours = int(input / 3600)
	minutes = int(math.Floor(float64(input%(3600)) / 60))
	return
}
