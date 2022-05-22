package widgets

import "math"

/*
var (
	name       nucular.TextEditor
	prevName   string
	nameResult []string
)

func init() {
	name.Flags = nucular.EditField
	name.Maxlen = 8
	name.SingleLine = true
}

func DrawItemFinder(w *nucular.Window, x, y int) (count int) {
	count = 6

	w.LayoutSpacePush(rect.Rect{
		X: x,
		Y: y,
		W: 120,
		H: 18,
	})
	w.Label("Find By Name:", "LC")
	y += 24

	w.LayoutSpacePush(rect.Rect{
		X: x,
		Y: y,
		W: 80,
		H: 22,
	})
	if e := name.Edit(w); e == nucular.EditActive || e == nucular.EditCommitted {
		l := len(name.Buffer)
		if l == 0 || l >= 2 {
			nameResult = nameResult[:0]
		}
		if l >= 2 {
			s := strings.ToLower(string(name.Buffer))
			if global.IsShowingPR() {
				for n, v := range pr.ItemsByName {
					if strings.Index(strings.ToLower(n), s) != -1 {
						nameResult = append(nameResult, fmt.Sprintf("%d - %s", v, n))
						count++
					}
				}
			} else {
				for n, v := range snes.ItemsByName {
					if strings.Index(strings.ToLower(n), s) != -1 {
						nameResult = append(nameResult, fmt.Sprintf("%s - %s", v, n))
						count++
					}
				}
			}
			sort.Strings(nameResult)
		}
	}
	w.LayoutSpacePush(rect.Rect{
		X: x + 90,
		Y: y,
		W: 40,
		H: 22,
	})
	y += 24

	for _, s := range nameResult {
		w.LayoutSpacePush(rect.Rect{
			X: x + 5,
			Y: y,
			W: 150,
			H: 22,
		})
		w.Label(s, "LC")
		y += 24
	}
	return
}
*/

func GetTime(input int) (hours int, minutes int) {
	hours = int(input / 3600)
	minutes = int(math.Floor(float64(input%(3600)) / 60))
	return
}
