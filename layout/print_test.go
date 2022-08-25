package layout

import "testing"

func TestPrint(t *testing.T) {
	height := 30
	width := 60
	setupBoard(height, width)
	makeTopAndBottom(0, 0, width-1, height-1)
	makeSides(1, 0, width-1, height-1)
	showBoard(height, width)
}
