package system

const BASE_POINTER = 0

var StackPtr = 0

func Push(Value int32) {
	if MEMORY != 0 && StackPtr < MEMORY {
		Stack[StackPtr] = Value
		StackPtr++
	}
}

func Pop() {
	if StackPtr != BASE_POINTER {
		StackPtr -= 1
		Stack[StackPtr] = 0
	} else {
		RaiseError("Value cannot be popped from empty stack.", StackPtr)
	}
}
