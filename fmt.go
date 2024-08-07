package errors

import (
	"fmt"
	"strconv"
	"unsafe"
)

// Format formats error message and adds location if present.
//
// fmt.Formatter interface implementation.
func (e wrapper) Format(s fmt.State, c rune) {
	e.formatMain(s, c)
	e.formatSub(s, c, true)
}

func (e wrapper) formatMain(s fmt.State, _ rune) {
	if e.msg == "" {
		e.msg = nomessage
	}

	fmt.Fprintf(s, "%s", e.msg)
}

func (e wrapper) formatSub(s fmt.State, c rune, delim bool) {
	if e.err == nil {
		return
	}

	if delim {
		if s.Flag(' ') {
			_, _ = s.Write([]byte{'\n'})
		} else {
			_, _ = s.Write([]byte{':', ' '})
		}
	}

	subPrintArg(s, e.err, c)
}

func (e withPC) Format(s fmt.State, c rune) {
	if e.msg != "" || s.Flag('+') || e.err == nil {
		e.wrapper.formatMain(s, c)
	}

	if e.pc != 0 && s.Flag('+') {
		if s.Flag(' ') {
			fmt.Fprintf(s, " at %+v", e.pc)
		} else {
			fmt.Fprintf(s, " (%v)", e.pc)
		}
	}

	e.wrapper.formatSub(s, c, e.msg != "" || s.Flag('+'))
}

func subFormat(s fmt.State, arg interface{}, verb rune) {
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
}

func bytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
