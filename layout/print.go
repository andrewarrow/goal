package layout

import "fmt"

func makeTopAndBottom(top, leading, width, height int) {
	board[top][leading] = "+"
	board[top][leading+width] = "+"
	for j := leading + 1; j < leading+width; j++ {
		board[top][j] = "-"
		board[height+top][j] = "-"
	}
	board[height+top][leading] = "+"
	board[height+top][leading+width] = "+"
}

func makeSides(top, leading, width, height int) {
	for i := top; i < height+top-1; i++ {
		board[i][leading] = "|"
		board[i][leading+width] = "|"
	}
}

func makeText(top, leading int, text string) {
	for i := 0; i < len(text); i++ {
		board[top][leading+i] = string(text[i])
	}
}

func addRenderedViewToBoard(rv *RenderedView) {
	makeSides(rv.Top+1, rv.Leading, rv.Width-1, rv.Height)
	makeTopAndBottom(rv.Top, rv.Leading, rv.Width-1, rv.Height-1)
}

func showBoard(rows, cols int) {
	//fmt.Println("   01234567890")
	for i := 0; i < rows; i++ {
		//fmt.Printf("%02d ", i)
		for j := 0; j < cols; j++ {
			fmt.Printf(stringCharAt(i, j))
		}
		if i < rows-1 {
			fmt.Printf("\n")
		}
	}
	fmt.Println("")
}
