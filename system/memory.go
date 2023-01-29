package system

import (
	"fmt"
	"unsafe"
)

const MEMORY = 300

var Heap [MEMORY / 2]Instruction
var Stack [MEMORY / 2]int32

func BlitMemory() {
	fmt.Println("Stack:\n", Stack)
	fmt.Println("Heap:\n", Heap)
	fmt.Println("Stack Size: ", unsafe.Sizeof(Stack))
	fmt.Println("Heap Size: ", unsafe.Sizeof(Heap))
}

func ClearAllMemory() {
	var EmptyInst Instruction
	var Counter = 0

	for Counter < len(Heap) {
		Heap[Counter] = EmptyInst
		Counter++
	}

	Counter = 0

	for Counter < len(Stack) {
		Stack[Counter] = 0
		Counter++
	}
}

func ClearHeap() {
	var EmptyInst Instruction
	var Counter = 0

	for Counter < len(Heap) {
		Heap[Counter] = EmptyInst
		Counter++
	}
}

func ClearStack() {
	var Counter = 0

	for Counter < len(Stack) {
		Stack[Counter] = 0
		Counter++
	}
}
