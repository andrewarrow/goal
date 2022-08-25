package layout

import (
	"encoding/json"
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
	showBoard(rows, cols)
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
		}

		if referencedViewLeading.renderedView.LeadingSet {
			view.renderedView.Leading = referencedViewLeading.renderedView.Leading + (view.Leading.Constant / 10)
			view.renderedView.LeadingSet = true
		}

		if view.Class == "UILabel" {
			view.renderedView.Width = len(view.Text)
			view.renderedView.Height = 1
			view.renderedView.WidthSet = true
			view.renderedView.HeightSet = true
		} else {
			if referencedViewTrailing.renderedView.WidthSet {
				view.renderedView.Width = computeWidth(view, referencedViewTrailing)
				view.renderedView.WidthSet = true
			}

			if referencedViewBottom.renderedView.HeightSet {
				view.renderedView.Height = computeHeight(view, referencedViewBottom)
				view.renderedView.HeightSet = true
			}
		}
	}
	for _, subview := range subviews {
		copyOfSubview := subview
		processSubviewsToRender(copyOfSubview, subview.Subviews)
	}
}

func computeWidth(view, referencedView *View) int {
	if view.Trailing.Constant < 0 {
		return referencedView.renderedView.Width - 6
	}
	return referencedView.renderedView.Width + 6
}

func computeHeight(view, referencedView *View) int {
	if view.Bottom.Constant < 0 {
		return referencedView.renderedView.Height - 4
	}
	return referencedView.renderedView.Height
}

func processSubviewsToPrint(view *View, subviews []*View) {
	if view.Id != "root" {
		renderedView := view.renderedView
		if view.Class == "UILabel" {
			makeText(renderedView.Top, renderedView.Leading, view.Text)
		} else {
			addRenderedViewToBoard(renderedView)
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
