package layout

import (
	"fmt"
	"testing"
)

func TestLayoutSetup(t *testing.T) {
	LoadFromFile("../examples/layout1.json")
	setupRootAndBoard(30, 60)
	processSubviewsForIdMap(&root.Root, root.Root.Subviews)
	views := []string{"root", "view1", "view2", "view3", "label1", "label2"}
	for _, viewName := range views {
		if idMap[viewName] == nil {
			t.Errorf("%s nil", viewName)
		}
	}
}

func TestLayoutN(t *testing.T) {
	LoadFromFile("../examples/layout1.json")
	setupRootAndBoard(30, 60)
	processSubviewsForIdMap(&root.Root, root.Root.Subviews)
	processSubviewsToRender(&root.Root, root.Root.Subviews)
	views := []string{"root", "view1", "view2", "view3", "label1", "label2"}
	for _, viewName := range views {
		view := idMap[viewName]
		fmt.Println(viewName, view.renderedView.String())
	}
}
