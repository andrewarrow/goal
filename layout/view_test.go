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

func TestLayoutRenders(t *testing.T) {
	LoadFromFile("../examples/layout1.json")
	setupRootAndBoard(30, 60)
	processSubviewsForIdMap(&root.Root, root.Root.Subviews)
	processSubviewsToRender(&root.Root, root.Root.Subviews)
	views := []string{"root", "view1", "view2", "view3", "label1", "label2"}
	for _, viewName := range views {
		view := idMap[viewName]
		fmt.Println(viewName, view.renderedView.String())
	}
	rootTop := idMap["root"].renderedView.Top == 0
	rootLeading := idMap["root"].renderedView.Leading == 0
	if !rootTop || !rootLeading {
		t.Errorf("root Top or Leading != 0")
	}
	rootWidth := idMap["root"].renderedView.Width == 60
	rootHeight := idMap["root"].renderedView.Height == 30
	if !rootWidth || !rootHeight {
		t.Errorf("root Width or Height wrong")
	}

	view1Top := idMap["view1"].renderedView.Top == 2
	view1Leading := idMap["view1"].renderedView.Leading == 3
	if !view1Top || !view1Leading {
		t.Errorf("view1 Top or Leading wrong")
	}
	view1Width := idMap["view1"].renderedView.Width == 60-6
	view1Height := idMap["view1"].renderedView.Height == 30-4
	if !view1Width || !view1Height {
		t.Errorf("view1 Width or Height wrong")
	}
}
