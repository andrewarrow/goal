package layout

import (
	"encoding/json"
	"fmt"
	"goal/files"
)

type Layout struct {
	Root View `json:"root"`
}

type View struct {
	Id       string     `json:"id"`
	Class    string     `json:"class"`
	Subviews []View     `json:"subviews"`
	Leading  Constraint `json:"leading"`
	Top      Constraint `json:"top"`
	Trailing Constraint `json:"trailing"`
	Bottom   Constraint `json:"bottom"`
	Text     string     `json:"text"`
}

type Constraint struct {
	Equal    string `json:"equal"`
	Constant int    `json:"constant"`
}

var root Layout

func LoadFromFile(filename string) {
	asString := files.ReadFile(filename)
	json.Unmarshal([]byte(asString), &root)
}

func Print(cols, rows int) {
	processSubviews(nil, &root.Root, root.Root.Subviews)

	//printTop(cols)
	//printRow(cols)
}

func processSubviews(superview, view *View, subviews []View) {
	if superview != nil {
		fmt.Println(superview.Id, view.Id, len(subviews))
	}
	if len(subviews) == 0 {
		// for now assume leaf is UILabel with text
		// text has a default width height based on size of font
		// in our case of CLI it's a fixed width font, so len of string is the width
		// and height is always 1 row
		fmt.Println("leaf", view.Text)
	}
	for _, subview := range subviews {
		copyOfSubview := subview
		processSubviews(view, &copyOfSubview, subview.Subviews)
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
