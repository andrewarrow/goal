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
	renderedView *RenderedView
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
	rootRenderedView.setComplete()
	root.Root.renderedView = &rootRenderedView
	processSubviewsForIdMap(&root.Root, root.Root.Subviews)

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
		processSubviewsToRender(&root.Root, root.Root.Subviews)
		if allViewsReady() {
			break
		}
	}
	processSubviewsToPrint(&root.Root, root.Root.Subviews)

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
		if view.renderedView.isComplete() == false {
			return false
		}
	}
	return true
}

func stringCharAt(row, col int) string {
	return charStringMaps[row][col]
}

func processSubviewsForIdMap(view *View, subviews []*View) {
	if view.Id != "root" {
		renderedView := RenderedView{}
		view.renderedView = &renderedView
	}
	idMap[view.Id] = view
	for _, subview := range subviews {
		copyOfSubview := subview
		processSubviewsForIdMap(copyOfSubview, subview.Subviews)
	}
}

func parseEqual(s string) (string, string) {
	tokens := strings.Split(s, ".")
	return tokens[0], tokens[1]
}

func processSubviewsToRender(view *View, subviews []*View) {
	if view.Id != "root" {
		id, position := parseEqual(view.Top.Equal)
		fmt.Println(id, position, view.Id, len(subviews))
		referencedView := idMap[id]

		if view.Class == "UILabel" {
			view.renderedView.Width = len(view.Text)
			view.renderedView.Height = 1
			view.renderedView.WidthSet = true
			view.renderedView.HeightSet = true
		}

		if referencedView.renderedView.TopSet {
			view.renderedView.Top = referencedView.renderedView.Top + 2
			view.renderedView.TopSet = true
		}

		if referencedView.renderedView.LeadingSet {
			view.renderedView.Leading = referencedView.renderedView.Leading + 3
			view.renderedView.LeadingSet = true
		}

		if view.Class != "UILabel" {
			if referencedView.renderedView.WidthSet {
				view.renderedView.Width = referencedView.renderedView.Width - 4
				view.renderedView.WidthSet = true
			}

			if referencedView.renderedView.HeightSet {
				view.renderedView.Height = referencedView.renderedView.Height - 4
				view.renderedView.HeightSet = true
			}
		}
	}
	for _, subview := range subviews {
		copyOfSubview := subview
		processSubviewsToRender(copyOfSubview, subview.Subviews)
	}
}

func processSubviewsToPrint(view *View, subviews []*View) {
	if view.Id != "root" {
		id, position := parseEqual(view.Top.Equal)
		fmt.Println("1", id, position, view.Id, len(subviews))
		renderedView := view.renderedView
		if view.Class == "UILabel" {
			makeText(renderedView.Top, renderedView.Leading, view.Text)
		} else {
			makeSides(renderedView.Top, renderedView.Leading, renderedView.Width, renderedView.Height+1)
			makeTopAndBottom(renderedView.Top, renderedView.Leading, renderedView.Width, renderedView.Height+1)
		}
	}
	for _, subview := range subviews {
		copyOfSubview := subview
		processSubviewsToPrint(copyOfSubview, subview.Subviews)
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
