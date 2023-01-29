package main

import (
	"fmt"
	"strings"
	"vm/system"
)

func main() {
	var ShellRunning = true

	for ShellRunning {
		var Command, Arg string

		fmt.Print("% ")
		fmt.Scanln(&Command, &Arg)

		if strings.HasPrefix(Command, "load") {
			system.ReadFile(Arg)
		} else if Command == "run" {
			system.Run()
		} else if Command == "mem" {
			system.BlitMemory()
		} else if Command == "clear_mem" {
			system.ClearAllMemory()
		} else if Command == "clear_heap" {
			system.ClearHeap()
		} else if Command == "clear_stack" {
			system.ClearHeap()
		} else if Command == "off" {
			var Option string

			fmt.Print("confirm[y/n]% ")
			fmt.Scan(&Option)

			if Option == "y" {
				ShellRunning = false
			}
		}
	}
}
