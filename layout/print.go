package layout

func makeTopAndBottom(row, leading, width, height int) {
	board[row][leading] = "+"
	board[row][leading+width] = "+"
	for j := leading + 1; j < leading+width; j++ {
		board[row][j] = "-"
		board[height][j] = "-"
	}
	board[height][leading] = "+"
	board[height][leading+width] = "+"
}

func makeSides(top, leading, width, height int) {
	for i := top; i < height; i++ {
		board[i][leading] = "|"
		board[i][leading+width] = "|"
	}
}

func makeText(top, leading int, text string) {
	for i := 0; i < len(text); i++ {
		board[top][leading+i] = string(text[i])
	}
}
