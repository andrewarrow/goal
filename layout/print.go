package layout

func makeTopAndBottom(row, leading, width, height int) {
	charStringMaps[row][leading] = "+"
	charStringMaps[row][width] = "+"
	for j := leading + 1; j < width; j++ {
		charStringMaps[row][j] = "-"
		charStringMaps[height][j] = "-"
	}
	charStringMaps[height][leading] = "+"
	charStringMaps[height][width] = "+"
}

func makeSides(top, leading, width, height int) {
	for i := top; i < height; i++ {
		charStringMaps[i][leading] = "|"
		charStringMaps[i][width] = "|"
	}
}
