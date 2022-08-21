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
	Id       string `json:"id"`
	Class    string `json:"class"`
	Subviews []View `json:"subviews"`
}

var root Layout

func LoadFromFile() {
	asString := files.ReadFile("layout.json")
	json.Unmarshal([]byte(asString), &root)
}

func Print() {
	processSubviews("root", root.Root.Subviews)
}

func processSubviews(id string, subviews []View) {
	fmt.Println(id, len(subviews))
	for _, subview := range subviews {
		processSubviews(subview.Id, subview.Subviews)
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
