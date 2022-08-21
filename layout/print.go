package layout

import (
	"fmt"
	"strings"
)

func printTop(cols int) {
	line := makeString(cols-2, "-")
	fmt.Printf("+%s+\n", line)
}

func printRow(cols int) {
	line := makeString(cols, " ")
	fmt.Printf("|%s|\n", line)
}

func makeString(cols int, s string) string {
	buff := []string{}
	for i := 0; i < cols; i++ {
		buff = append(buff, s)
	}
	return strings.Join(buff, "")
}
