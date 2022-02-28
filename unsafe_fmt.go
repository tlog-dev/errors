package errors

import (
	"fmt"
	"io"
	"strconv"
	"unsafe"
)

func subPrintArg(s fmt.State, arg interface{}, verb rune) {
	i := *(*iface)(unsafe.Pointer(&s))
	if i.typ != ppType {
		var buf [64]byte

		i := 0

		buf[i] = '%'
		i++

		for _, f := range "-+# 0" {
			if s.Flag(int(f)) {
				buf[i] = byte(f)
				i++
			}
		}

		if w, ok := s.Width(); ok {
			q := strconv.AppendInt(buf[:i], int64(w), 10)
			i = len(q)
		}

		if p, ok := s.Precision(); ok {
			buf[i] = '.'
			i++

			q := strconv.AppendInt(buf[:i], int64(p), 10)
			i = len(q)
		}

		buf[i] = byte(verb)
		i++

		fmt.Fprintf(s, bytesToString(buf[:i]), arg)

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

func (formatter) Format(s fmt.State, c rune) {
	i := *(*iface)(unsafe.Pointer(&s))

	ppType = i.typ
}

func bytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
