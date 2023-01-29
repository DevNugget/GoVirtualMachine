package system

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadFile(FileName string) {
	start := time.Now()
	var LineCount = 0

	File, Error := os.Open(FileName)

	if Error != nil {
		fmt.Println(Error)
	}

	FileScanner := bufio.NewScanner(File)
	FileScanner.Split(bufio.ScanLines)

	for FileScanner.Scan() {
		var Inst Instruction
		var Arg = strings.Split(FileScanner.Text(), " ")
		var Opcode, _ = strconv.ParseInt(Arg[0], 10, 0)

		if len(Arg) > 1 {
			var ArgCounter = 2

			for ArgCounter < len(Arg) {
				var Val, _ = strconv.ParseInt(Arg[ArgCounter], 10, 0)
				Inst.Args[ArgCounter-2] = int32(Val)
				ArgCounter++
			}
		}

		Inst.Name = Arg[1]
		Inst.Opcode = int32(Opcode)

		Heap[Opcode] = Inst

		LineCount++
	}

	File.Close()
	fmt.Println("Instructions loaded to Heap memory.")
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
