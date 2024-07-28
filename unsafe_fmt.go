//go:build ignore
// +build ignore

package errors

import (
	"fmt"
	"io"
	"unsafe"
)

func subPrintArg(s fmt.State, arg interface{}, verb rune) {
	i := *(*iface)(unsafe.Pointer(&s))
	if i.typ != ppType {
		subFormat(s, arg, verb)
		return
	}

	printArg(i.word, arg, verb)
}

type (
	iface struct {
		typ, word unsafe.Pointer
	}

	formatter struct{}
)

var ppType unsafe.Pointer

func init() {
	fmt.Fprintf(io.Discard, "%v", formatter{})
}

//go:linkname printArg fmt.(*pp).printArg
func printArg(p unsafe.Pointer, arg interface{}, verb rune)

//go:linkname newPrinter fmt.newPrinter
func newPrinter() unsafe.Pointer

//go:linkname ppFree fmt.(*pp).free
func ppFree(unsafe.Pointer)

func (formatter) Format(s fmt.State, _ rune) {
	i := *(*iface)(unsafe.Pointer(&s))

	ppType = i.typ
}
