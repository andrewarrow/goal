package main

import "fmt"

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  goal run                       # use example/layout.json with cols=60 and rows=30")
	fmt.Println("  goal run ~/layout.json 80 40   # use passed in file, cols, and rows")
	fmt.Println("")
}
