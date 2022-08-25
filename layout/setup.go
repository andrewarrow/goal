package layout

func setupRootAndBoard(rows, cols int) *RenderedView {
	rootRenderedView := RenderedView{}
	rootRenderedView.Width = cols
	rootRenderedView.Height = rows
	rootRenderedView.setComplete()
	root.Root.renderedView = &rootRenderedView
	for i := 0; i < rows; i++ {
		board[i] = map[int]string{}
		for j := 0; j < cols; j++ {
			board[i][j] = " "
		}
	}

	return &rootRenderedView
}
