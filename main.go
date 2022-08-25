package main

import (
	"goal/layout"
	"math/rand"
	"os"
	"strconv"
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
		filename := "examples/layout1.json"
		cols := 60
		rows := 30
		if len(os.Args) > 4 {
			filename = os.Args[2]
			cols, _ = strconv.Atoi(os.Args[3])
			rows, _ = strconv.Atoi(os.Args[4])
		}
		layout.LoadFromFile(filename)
		layout.Print(cols, rows)
	} else if command == "add" {
	}
}
