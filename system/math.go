package system

func StackAdd() {
	var Operand1 = Stack[StackPtr-1]
	Pop()
	var Operand2 = Stack[StackPtr-1]
	Pop()
	Push(Operand1 + Operand2)
}
