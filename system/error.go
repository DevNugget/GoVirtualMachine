package system

import (
	"fmt"
	"os"
)

func RaiseError(Msg string, Pointer int) {
	fmt.Println(Msg, '\n', Pointer)
	os.Exit(1)
}
