package errors

import (
	"fmt"
	"os"
	"reflect"
	"unsafe"
)

func subPrintArg(s fmt.State, arg interface{}, verb rune) {
	if tp := reflect.TypeOf(s); tp.String() != "*fmt.pp" {
		fmt.Fprintf(os.Stderr, "aborting subprint (state type: %v)\n", tp.String())
		return
	}

	type iface struct {
		typ, word unsafe.Pointer
	}

	i := *(*iface)(unsafe.Pointer(&s))

	printArg(i.word, arg, verb)
}

//go:linkname printArg fmt.(*pp).printArg
func printArg(p unsafe.Pointer, arg interface{}, verb rune)
