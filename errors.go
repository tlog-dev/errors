package errors

import (
	"fmt"
	"runtime"

	"github.com/nikandfor/tlog"
)

type (
	Location = tlog.Location

	wrapper struct {
		err error
		msg string
		loc Location
	}
)

func (e wrapper) Error() string {
	if e.err == nil {
		return e.msg
	}
	if e.msg == "" {
		return e.err.Error()
	}
	return e.msg + ": " + e.err.Error()
}

func New(f string, args ...interface{}) error {
	return wrapper{
		msg: fmt.Sprintf(f, args...),
		loc: Caller(1),
	}
}

func NewNoLoc(f string, args ...interface{}) error {
	return wrapper{
		msg: fmt.Sprintf(f, args...),
	}
}

func Wrap(err error, f string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return wrapper{
		err: err,
		msg: fmt.Sprintf(f, args...),
		loc: Caller(1),
	}
}

func WrapNoLoc(err error, f string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return wrapper{
		err: err,
		msg: fmt.Sprintf(f, args...),
	}
}

func Unwrap(err error) error {
	switch e := err.(type) {
	case wrapper:
		return e.err
	case interface{ Unwrap() error }:
		return e.Unwrap()
	default:
		return nil
	}
}

// Caller returns information about the calling goroutine's stack. The argument s is the number of frames to ascend, with 0 identifying the caller of Caller.
func Caller(s int) Location {
	var pc [1]uintptr
	runtime.Callers(2+s, pc[:])
	return Location(pc[0])
}

// Funcentry returns information about the calling goroutine's stack. The argument s is the number of frames to ascend, with 0 identifying the caller of Caller.
func Funcentry(s int) Location {
	var pc [1]uintptr
	runtime.Callers(2+s, pc[:])
	return Location(Location(pc[0]).Entry())
}
