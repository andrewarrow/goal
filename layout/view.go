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
var board = map[int]map[int]string{}

func LoadFromFile(filename string) {
	asString := files.ReadFile(filename)
	json.Unmarshal([]byte(asString), &root)
}

func Print(cols, rows int) {
	rootRenderedView := setupRootAndBoard(rows, cols)
	processSubviewsForIdMap(&root.Root, root.Root.Subviews)

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
	return board[row][col]
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
	if s == "" {
		return "", ""
	}
	tokens := strings.Split(s, ".")
	return tokens[0], tokens[1]
}

func processSubviewsToRender(view *View, subviews []*View) {
	if view.Id != "root" {
		topId, _ := parseEqual(view.Top.Equal)
		leadingId, _ := parseEqual(view.Leading.Equal)
		trailingId, _ := parseEqual(view.Trailing.Equal)
		bottomId, _ := parseEqual(view.Bottom.Equal)

		referencedViewTop := idMap[topId]
		referencedViewLeading := idMap[leadingId]
		referencedViewTrailing := idMap[trailingId]
		referencedViewBottom := idMap[bottomId]

		if referencedViewTop.renderedView.TopSet {
			view.renderedView.Top = referencedViewTop.renderedView.Top + 2
			view.renderedView.TopSet = true
			//fmt.Println("|", view.Id, "topSet", view.renderedView.Top)
		}

		if referencedViewLeading.renderedView.LeadingSet {
			view.renderedView.Leading = referencedViewLeading.renderedView.Leading + (view.Leading.Constant / 10)
			view.renderedView.LeadingSet = true
			//fmt.Println("|", view.Id, "leadingSet", view.renderedView.Leading)
		}

		if view.Class == "UILabel" {
			view.renderedView.Width = len(view.Text)
			view.renderedView.Height = 1
			view.renderedView.WidthSet = true
			view.renderedView.HeightSet = true
		} else {
			// root.leading to root.trailing == root.width
			// view1.leading to label1.trailing == label1.width
			// label1.leading + 99 to label2.trailing == label2.width
			if referencedViewTrailing.renderedView.WidthSet {
				//fmt.Println("w", view.Id, referencedViewTrailing.Id, referencedViewTrailing.renderedView.Width)
				view.renderedView.Width = 24 //referencedViewTrailing.renderedView.Width - 4
				view.renderedView.WidthSet = true
			}

			if referencedViewBottom.renderedView.HeightSet {
				//fmt.Println("h", view.Id, referencedViewBottom.Id, referencedViewBottom.renderedView.Height)
				view.renderedView.Height = 24 //referencedViewBottom.renderedView.Height - 4
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
