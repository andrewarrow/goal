package layout

import "testing"

func TestPrint(t *testing.T) {
	height := 10
	width := 20
	setupBoard(height, width)

	rv := RenderedView{}
	rv.Top = 0
	rv.Leading = 0
	rv.Width = width
	rv.Height = height

	addRenderedViewToBoard(&rv)
	showBoard(height, width)
}
