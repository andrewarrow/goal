package layout

import "fmt"

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

func showBoard(rows, cols int) {
	for i := 0; i < rows; i++ {
		fmt.Printf("%02d ", i)
		for j := 0; j < cols; j++ {
			fmt.Printf(stringCharAt(i, j))
		}
		if i < rows-1 {
			fmt.Printf("\n")
		}
	}
	fmt.Println("")
}
