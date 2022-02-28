package errors

import (
	"fmt"

	"github.com/nikandfor/loc"
)

type (
	// PC is a program counter and represents location in a source code.
	PC  = loc.PC
	PCs = loc.PCs

	wrapper struct {
		err error
		msg string
	}

	withPC struct {
		wrapper
		pc PC
	}

	withPCs struct {
		wrapper
		pcs PCs
	}

	Locationer interface {
		Location() PC
	}
)

const nomessage = "(no message)"

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(f string, args ...interface{}) error {
	return withPC{
		wrapper: wrapper{
			msg: fmt.Sprintf(f, args...),
		},
		pc: loc.Caller(1),
	}
}

// NewNoLoc is like a New but with no caller info.
func NewNoLoc(f string, args ...interface{}) error {
	return wrapper{
		msg: fmt.Sprintf(f, args...),
	}
}

// NewDepth returns an error that formats as the given text.
// Location where error was created (d frames higher) is recorded.
// Each call to New returns a distinct error value even if the text is identical.
func NewDepth(d int, f string, args ...interface{}) error {
	return withPC{
		wrapper: wrapper{
			msg: fmt.Sprintf(f, args...),
		},
		pc: loc.Caller(d + 1),
	}
}

// NewLoc returns an error with given PC that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func NewLoc(pc PC, f string, args ...interface{}) error {
	return withPC{
		wrapper: wrapper{
			msg: fmt.Sprintf(f, args...),
		},
		pc: pc,
	}
}

// Wrap returns an error that describes given error with given text.
// Returns nil if err is nil.
func Wrap(err error, f string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return withPC{
		wrapper: wrapper{
			err: err,
			msg: fmt.Sprintf(f, args...),
		},
		pc: loc.Caller(1),
	}
}

// WrapNoLoc is like Wrap but without caller info.
func WrapNoLoc(err error, f string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return wrapper{
		err: err,
		msg: fmt.Sprintf(f, args...),
	}
}

// WrapDepth returns an error that describes given error with given text.
// Location where error was created (d frames higher) is recorded.
// Returns nil if err is nil.
func WrapDepth(err error, d int, f string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return withPC{
		wrapper: wrapper{
			err: err,
			msg: fmt.Sprintf(f, args...),
		},
		pc: loc.Caller(d + 1),
	}
}

// WrapLoc returns an error with given PC that describes given error with given text.
// Returns nil if err is nil.
func WrapLoc(err error, pc PC, f string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return withPC{
		wrapper: wrapper{
			err: err,
			msg: fmt.Sprintf(f, args...),
		},
		pc: pc,
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

// PC returns underlaying error location.
func (e withPC) Location() PC {
	return e.pc
}
