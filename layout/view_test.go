package layout

import (
	"fmt"
	"testing"
)

func TestLayout(t *testing.T) {
	LoadFromFile("../examples/layout1.json")
	setupRootAndBoard(60, 30)
	processSubviewsForIdMap(&root.Root, root.Root.Subviews)
	fmt.Println(idMap)
}
