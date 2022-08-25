package layout

import (
	"testing"
)

func TestLayoutSetup(t *testing.T) {
	LoadFromFile("../examples/layout1.json")
	setupRootAndBoard(30, 60)
	processSubviewsForIdMap(&root.Root, root.Root.Subviews)
	if idMap["root"] == nil {
		t.Errorf("root nil")
	}
	if idMap["view1"] == nil {
		t.Errorf("view1 nil")
	}
	if idMap["view2"] == nil {
		t.Errorf("view1 nil")
	}
	if idMap["view3"] == nil {
		t.Errorf("view1 nil")
	}
	if idMap["label1"] == nil {
		t.Errorf("label1 nil")
	}
	if idMap["label2"] == nil {
		t.Errorf("label2 nil")
	}
}
