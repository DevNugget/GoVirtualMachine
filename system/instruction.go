package system

type Instruction struct {
	Name   string
	Opcode int32
	Args   [3]int32
}

func ReadInstruction(Inst Instruction) {
	if Inst.Name == "push" {
		Push(Inst.Args[0])
	} else if Inst.Name == "pop" {
		Pop()
	} else if Inst.Name == "add" {
		StackAdd()
	}
}
