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
{
  "root": {
    "view1": {
      "label1": "foo1"
    },
    "view2": {
      "label1": "foo2",
      "label2": "foo3"
    }
  }
}


+----------------------------+
|                            |
|  foo1  foo2foo3            |
|                            |
+----------------------------+

{
  "root": {
    "view1": {
			"view2": {
	  		"label1": "foo1"
			},
    },
  }
}
+----------------------------+
|                            |
|  +----------------------+  |
|  |                      |  |
|  |                      |  |
|  +----------------------+  |
|                            |
+----------------------------+

*/
