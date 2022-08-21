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

func LoadFromFile() {
	asString := files.ReadFile("layout.json")
	json.Unmarshal([]byte(asString), &root)
}

func Print() {
	processSubviews(root.Root, root.Root.Subviews)
}

func processSubviews(view View, subviews []View) {
	fmt.Println(view, len(subviews))
	for _, subview := range subviews {
		processSubviews(subview, subview.Subviews)
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
