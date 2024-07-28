package errors

import "fmt"

func subPrintArg(s fmt.State, arg interface{}, verb rune) {
	subFormat(s, arg, verb)
}
