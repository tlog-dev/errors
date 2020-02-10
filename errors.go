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

const nomessage = "(no message)"

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(f string, args ...interface{}) error {
	return wrapper{
		msg: fmt.Sprintf(f, args...),
	}
}

// NewHere returns an error that formats as the given text.
// Location where error was created is recorded.
// Each call to New returns a distinct error value even if the text is identical.
func NewHere(f string, args ...interface{}) error {
	return wrapper{
		msg: fmt.Sprintf(f, args...),
		loc: Caller(1),
	}
}

// NewDepth returns an error that formats as the given text.
// Location where error was created (d frames higher) is recorded.
// Each call to New returns a distinct error value even if the text is identical.
func NewDepth(d int, f string, args ...interface{}) error {
	return wrapper{
		msg: fmt.Sprintf(f, args...),
		loc: Caller(d + 1),
	}
}

// NewLoc returns an error with given Location that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func NewLoc(loc Location, f string, args ...interface{}) error {
	return wrapper{
		msg: fmt.Sprintf(f, args...),
		loc: loc,
	}
}

// New returns an error that describes given error with given text.
// Returns nil if err is nil.
func Wrap(err error, f string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return wrapper{
		err: err,
		msg: fmt.Sprintf(f, args...),
	}
}

// WrapHere returns an error that describes given error with given text.
// Location where error was wrapped is recorded.
// Returns nil if err is nil.
func WrapHere(err error, f string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return wrapper{
		err: err,
		msg: fmt.Sprintf(f, args...),
		loc: Caller(1),
	}
}

// WrapDepth returns an error that describes given error with given text.
// Location where error was created (d frames higher) is recorded.
// Returns nil if err is nil.
func WrapDepth(d int, err error, f string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return wrapper{
		err: err,
		msg: fmt.Sprintf(f, args...),
		loc: Caller(d + 1),
	}
}

// WrapLoc returns an error with given Location that describes given error with given text.
// Returns nil if err is nil.
func WrapLoc(loc Location, err error, f string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return wrapper{
		err: err,
		msg: fmt.Sprintf(f, args...),
		loc: loc,
	}
}

// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
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

// Error is an error interface implementation.
func (e wrapper) Error() string {
	if e.err == nil {
		if e.msg == "" {
			return nomessage
		}
		return e.msg
	}
	if e.msg == "" {
		return e.err.Error()
	}
	return e.msg + ": " + e.err.Error()
}

// Unwrap returns original error that was wrapped or nil if it's not wrapper.
func (e wrapper) Unwrap() error {
	return e.err
}

// Location returns underlaying error location.
func (e wrapper) Location() Location {
	return e.loc
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
