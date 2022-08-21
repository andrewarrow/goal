package files

import (
	"fmt"
	"io/ioutil"
)

func ReadFile(name string) string {
	fname := fmt.Sprintf("data/%s", name)
	b, _ := ioutil.ReadFile(fname)
	return string(b)
}
