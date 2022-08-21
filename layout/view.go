package layout

import (
	"encoding/json"
	"fmt"
	"goal/files"
)

type Layout struct {
}

func LoadFromFile() {
	asString := files.ReadFile("layout.json")
	var asMap map[string]any
	json.Unmarshal([]byte(asString), &asMap)
	fmt.Println(asMap)
}
