package main

import (
	"goal/layout"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "run" {
		layout.LoadFromFile()
		layout.Print()
	} else if command == "add" {
	}
}
