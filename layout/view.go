package layout

import (
	"encoding/json"
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
	root := rootMap["root"]
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
