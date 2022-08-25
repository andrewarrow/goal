package layout

func setupRootAndBoard(rows, cols int) *RenderedView {
	rootRenderedView := RenderedView{}
	rootRenderedView.Width = cols
	rootRenderedView.Height = rows
	rootRenderedView.setComplete()
	root.Root.renderedView = &rootRenderedView
	setupBoard(rows, cols)

	return &rootRenderedView
}

func setupBoard(rows, cols int) {
	for i := 0; i < rows; i++ {
		board[i] = map[int]string{}
		for j := 0; j < cols; j++ {
			board[i][j] = " "
		}
	}
}
