package system

import (
	"fmt"
	"time"
)

func Run() {
	start := time.Now()

	var InstPtr = 0

	for InstPtr < len(Heap) {
		ReadInstruction(Heap[InstPtr])
		InstPtr++
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
