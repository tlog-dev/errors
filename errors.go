package errors

import (
	"fmt"
	"runtime"

	"github.com/nikandfor/tlog"
)

type (
	Location = tlog.Location

	wrapper struct {
		orig error
		msg  string
		loc  Location
	}

	causer interface {
		Cause() error
	}
)

func (e *wrapper) Error() string {
	if e.orig == nil {
		return e.msg
	}
	if e.msg == "" {
		return e.orig.Error()
	}
	return e.msg + ": " + e.orig.Error()
}

func New(msg string) *wrapper {
	return &wrapper{
		msg: msg,
		loc: Caller(1),
	}
}

func Newf(f string, args ...interface{}) error {
	return &wrapper{
		msg: fmt.Sprintf(f, args...),
		loc: Caller(1),
	}
}

func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	return &wrapper{
		orig: err,
		msg:  msg,
		loc:  Caller(1),
	}
}

func Wrapf(err error, f string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &wrapper{
		orig: err,
		msg:  fmt.Sprintf(f, args...),
		loc:  Caller(1),
	}
}

func WrapfNoLoc(err error, f string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &wrapper{
		orig: err,
		msg:  fmt.Sprintf(f, args...),
	}
}

func Cause(err error) error {
	for err != nil {
		switch e := err.(type) {
		case *wrapper:
			if e.orig == nil {
				return err
			}
			err = e.orig
		case causer:
			err = e.Cause()
		default:
			return err
		}
	}
	return nil
}

// Caller returns information about the calling goroutine's stack. The argument s is the number of frames to ascend, with 0 identifying the caller of Caller.
//
// It's hacked version of runtime.Caller with no allocs.
func Caller(s int) Location {
	var pc [1]uintptr
	runtime.Callers(2+s, pc[:])
	return Location(pc[0])
}

// Funcentry returns information about the calling goroutine's stack. The argument s is the number of frames to ascend, with 0 identifying the caller of Caller.
//
// It's hacked version of runtime.Callers -> runtime.CallersFrames -> Frames.Next -> Frame.Entry with no allocs.
func Funcentry(s int) Location {
	var pc [1]uintptr
	runtime.Callers(2+s, pc[:])
	return Location(Location(pc[0]).Entry())
}
