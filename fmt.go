package errors

import "fmt"

func (e wrapper) Format(s fmt.State, c rune) {
	if !s.Flag('+') {
		_, _ = s.Write([]byte(e.Error()))
		return
	}

	if e.loc == 0 {
		if e.err == nil {
			_, _ = s.Write([]byte(e.msg))
			return
		}

		var f string
		if s.Flag(' ') {
			f = "% +v"
		} else {
			f = "%+v"
		}

		if e.msg == "" {
			fmt.Fprintf(s, f, e.err)
		} else {
			fmt.Fprintf(s, e.msg+": "+f, e.err)
		}

		return
	}

	if s.Flag(' ') {
		_, file, line := e.loc.NameFileLine()
		switch {
		case e.err == nil:
			fmt.Fprintf(s, "%s\nat %v:%d", e.msg, file, line)
		case e.msg == "":
			fmt.Fprintf(s, "at %v:%d\n% +v", file, line, e.err)
		default:
			fmt.Fprintf(s, "%s\nat %v:%d\n% +v", e.msg, file, line, e.err)
		}
		return
	}

	switch {
	case e.err == nil:
		fmt.Fprintf(s, "%s (%v)", e.msg, e.loc)
	case e.msg == "":
		fmt.Fprintf(s, "(%v): %+v", e.loc, e.err)
	default:
		fmt.Fprintf(s, "%s (%v): %+v", e.msg, e.loc, e.err)
	}
}
