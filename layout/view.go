package layout

import (
	"encoding/json"
	"fmt"
	"goal/files"
)

type Layout struct {
}

var rootMap map[string]any

func LoadFromFile() {
	asString := files.ReadFile("layout.json")
	json.Unmarshal([]byte(asString), &rootMap)
}

func Print() {
	root := rootMap["root"].(map[string]any)
	processSubviews("root", root["subviews"].([]any))
}

func processSubviews(id string, subviews []any) {
	fmt.Println(id, len(subviews))
	for _, subview := range subviews {
		m := subview.(map[string]any)
		if m["subviews"] == nil {
			continue
		}
		processSubviews(m["id"].(string), m["subviews"].([]any))
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
