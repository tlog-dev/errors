package errors

import (
	"fmt"
)

// Format formats error message and adds location if present.
//
// fmt.Formatter interface implementation.
func (e wrapper) Format(s fmt.State, c rune) {
	e.formatMain(s, c)
	e.formatSub(s, c, true)
}

func (e wrapper) formatMain(s fmt.State, c rune) {
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
