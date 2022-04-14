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

	Callerer interface {
		Caller() PC
	}

	Callerser interface {
		Callers() PCs
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

// NewNoCaller is like a New but with no caller info.
func NewNoCaller(f string, args ...interface{}) error {
	return wrapper{
		msg: fmt.Sprintf(f, args...),
	}
}

// NewDepth returns an error that formats as the given text.
// Callsite where error was created (d frames higher) is recorded.
// Each call to New returns a distinct error value even if the text is identical.
func NewDepth(d int, f string, args ...interface{}) error {
	return withPC{
		wrapper: wrapper{
			msg: fmt.Sprintf(f, args...),
		},
		pc: loc.Caller(d + 1),
	}
}

// NewStack returns an error with message formatted in fmt package style.
// Caller frames are recorded (skipping d frames).
// Experimental, may be deleted at any time.
func NewStack(skip, n int, f string, args ...interface{}) error {
	return withPCs{
		wrapper: wrapper{
			msg: fmt.Sprintf(f, args...),
		},
		pcs: loc.Callers(skip+1, n),
	}
}

// NewCaller returns an error with given PC that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func NewCaller(pc PC, f string, args ...interface{}) error {
	return withPC{
		wrapper: wrapper{
			msg: fmt.Sprintf(f, args...),
		},
		pc: pc,
	}
}

// NewCallers returns an error with given PC that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
// Experimental, may be deleted at any time.
func NewCallers(pcs PCs, f string, args ...interface{}) error {
	return withPCs{
		wrapper: wrapper{
			msg: fmt.Sprintf(f, args...),
		},
		pcs: pcs,
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

// WrapNoCaller is like Wrap but without caller info.
func WrapNoCaller(err error, f string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return wrapper{
		err: err,
		msg: fmt.Sprintf(f, args...),
	}
}

// WrapDepth returns an error that describes given error with given text.
// Callsite where error was created (d frames higher) is recorded.
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

// Experimental, may be deleted at any time.
func WrapStack(err error, skip, n int, f string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return withPCs{
		wrapper: wrapper{
			err: err,
			msg: fmt.Sprintf(f, args...),
		},
		pcs: loc.Callers(skip+1, n),
	}
}

// WrapCaller returns an error with given PC that describes given error with given text.
// Returns nil if err is nil.
func WrapCaller(err error, pc PC, f string, args ...interface{}) error {
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

// Experimental, may be deleted at any time.
func WrapCallers(err error, pcs PCs, f string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return withPCs{
		wrapper: wrapper{
			err: err,
			msg: fmt.Sprintf(f, args...),
		},
		pcs: pcs,
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

// Caller returns underlaying error location.
func (e withPC) Caller() PC {
	return e.pc
}

// Caller returns underlaying error location.
func (e withPCs) Caller() PC {
	if len(e.pcs) == 0 {
		return 0
	}

	return e.pcs[0]
}

// Callers returns underlaying error location.
func (e withPCs) Callers() PCs {
	return e.pcs
}
