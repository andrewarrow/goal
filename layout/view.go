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
    "subviews": [
		  {
			  "name": "view1",
				"class": "UIView",
				"subviews": [
					{
						"name": "view2",
						"class": "UIView",
						"subviews": [
									{
										"name": "label1",
										"class": "UILabel",
										"text": "Hello"
									}
						]
					},
					{
						"name": "view3",
						"class": "UIView",
						"subviews": [
									{
										"name": "label2",
										"class": "UILabel",
										"text": "There"
									}
						]
					}
				]
      }
		]
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
