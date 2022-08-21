package layout

import (
	"fmt"
	"strings"
)

func printTop(cols int) {
	line := makeDashes(cols)
	fmt.Printf("+%s+\n", line)
}

func makeDashes(cols int) string {
	buff := []string{}
	for i := 0; i < cols; i++ {
		buff = append(buff, "-")
	}
	return strings.Join(buff, "")
}
