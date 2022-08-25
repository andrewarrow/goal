package layout

import "testing"

func TestPrint(t *testing.T) {
	height := 20
	width := 30
	setupBoard(height, width)

	rv := RenderedView{}
	rv.Top = 0
	rv.Leading = 0
	rv.Width = width
	rv.Height = height

	addRenderedViewToBoard(&rv)

	rv.Top = 2
	rv.Leading = 2
	rv.Width = 18
	rv.Height = 16

	addRenderedViewToBoard(&rv)

	rv.Top = 4
	rv.Leading = 4
	rv.Width = 8
	rv.Height = 6

	addRenderedViewToBoard(&rv)

	showBoard(height, width)
}
