package layout

import (
	"encoding/json"
	"fmt"
	"goal/files"
	"strings"
)

type Layout struct {
	Root View `json:"root"`
}

type View struct {
	Id           string     `json:"id"`
	Class        string     `json:"class"`
	Subviews     []*View    `json:"subviews"`
	Leading      Constraint `json:"leading"`
	Top          Constraint `json:"top"`
	Trailing     Constraint `json:"trailing"`
	Bottom       Constraint `json:"bottom"`
	Text         string     `json:"text"`
	RenderedView *RenderedView
}

type Constraint struct {
	Equal    string `json:"equal"`
	Constant int    `json:"constant"`
}

var root Layout
var idMap = map[string]*View{}
var charStringMaps = map[int]map[int]string{}

func LoadFromFile(filename string) {
	asString := files.ReadFile(filename)
	json.Unmarshal([]byte(asString), &root)
}

func Print(cols, rows int) {
	rootRenderedView := RenderedView{}
	rootRenderedView.Width = cols
	rootRenderedView.Height = rows
	rootRenderedView.SetComplete()
	root.Root.RenderedView = &rootRenderedView
	processSubviewsForIdMap(nil, &root.Root, root.Root.Subviews)

	fmt.Println(idMap)

	for i := 0; i < rows; i++ {
		charStringMaps[i] = map[int]string{}
		for j := 0; j < cols; j++ {
			charStringMaps[i][j] = " "
		}
	}

	makeTopAndBottom(0, 0, rootRenderedView.Width-1, rootRenderedView.Height-1)
	makeSides(1, 0, rootRenderedView.Width-1, rootRenderedView.Height-1)

	for {
		processSubviewsToRender(nil, &root.Root, root.Root.Subviews)
		if allViewsReady() {
			break
		}
	}
	processSubviewsToPrint(nil, &root.Root, root.Root.Subviews)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf(stringCharAt(i, j))
		}
		if i < rows-1 {
			fmt.Printf("\n")
		}
	}
	fmt.Println("")
}

func allViewsReady() bool {
	for _, view := range idMap {
		if view.RenderedView.isComplete() == false {
			return false
		}
	}
	return true
}

func stringCharAt(row, col int) string {
	return charStringMaps[row][col]
}

func processSubviewsForIdMap(superview, view *View, subviews []*View) {
	if superview != nil {
		fmt.Println(superview.Id, view.Id, len(subviews))
	}
	renderedView := RenderedView{}
	view.RenderedView = &renderedView
	idMap[view.Id] = view
	if len(subviews) == 0 {
		// for now assume leaf is UILabel with text
		// text has a default width height based on size of font
		// in our case of CLI it's a fixed width font, so len of string is the width
		// and height is always 1 row
		fmt.Println("leaf1", view.Text)
	}
	for _, subview := range subviews {
		copyOfSubview := subview
		processSubviewsForIdMap(view, copyOfSubview, subview.Subviews)
	}
}

func parseEqual(s string) (string, string) {
	tokens := strings.Split(s, ".")
	return tokens[0], tokens[1]
}

func processSubviewsToRender(superview, view *View, subviews []*View) {
	if superview != nil {
		id, position := parseEqual(view.Top.Equal)
		fmt.Println(id, position, superview.Id, view.Id, len(subviews))
		renderedView := RenderedView{}
		renderedView.Top = 2
		renderedView.Leading = 3
		renderedView.Width = superview.RenderedView.Width - 4
		renderedView.Height = superview.RenderedView.Height - 4
		view.RenderedView = &renderedView
	}
	if len(subviews) == 0 {
		fmt.Println("leaf2", view.Text)
	}
	for _, subview := range subviews {
		copyOfSubview := subview
		processSubviewsToRender(view, copyOfSubview, subview.Subviews)
	}
}

func processSubviewsToPrint(superview, view *View, subviews []*View) {
	if superview != nil {
		id, position := parseEqual(view.Top.Equal)
		fmt.Println("1", id, position, superview.Id, view.Id, len(subviews))
		renderedView := view.RenderedView
		if view.Class == "UILabel" {
			makeText(renderedView.Top, renderedView.Leading, view.Text)
		} else {
			makeSides(renderedView.Top, renderedView.Leading, renderedView.Width, renderedView.Height+1)
			makeTopAndBottom(renderedView.Top, renderedView.Leading, renderedView.Width, renderedView.Height+1)
		}
	}
	if len(subviews) == 0 {
		fmt.Println("leaf3", view.Text)
	}
	for _, subview := range subviews {
		copyOfSubview := subview
		processSubviewsToPrint(view, copyOfSubview, subview.Subviews)
	}
}

/*
+---------------------------------------------------+
|                                                   |
|  +---------------------------------------------+  |
|  |                                             |  |
|  |  +---------+  +---------+                   |  |
|  |  |         |  |         |                   |  |
|  |  |  Hello  |  |  There  |                   |  |
|  |  |         |  |         |                   |  |
|  |  |         |  |         |                   |  |
|  |  |         |  |         |                   |  |
|  |  |         |  |         |                   |  |
|  |  |         |  |         |                   |  |
|  |  |         |  |         |                   |  |
|  |  +---------+  +---------+                   |  |
|  |                                             |  |
|  +---------------------------------------------+  |
|                                                   |
+---------------------------------------------------+
*/
