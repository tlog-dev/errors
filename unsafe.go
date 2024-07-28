//go:build go1.13 && ignore
// +build go1.13,ignore

//nolint:godot
package errors

//nolint:gci
import (
	_ "errors" // for go:linkname
	_ "unsafe" // for go:linkname
)

//go:linkname Is errors.Is

// Is reports whether any error in err's chain matches target.
//
// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
//
// An error is considered to match a target if it is equal to that target or if
// it implements a method Is(error) bool such that Is(target) returns true.
//
// link to stdlib errors.Is
func Is(err, target error) bool

//go:linkname As errors.As

// As finds the first error in err's chain that matches target, and if so, sets
// target to that error value and returns true.
//
// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
//
// An error matches target if the error's concrete value is assignable to the value
// pointed to by target, or if the error has a method As(interface{}) bool such that
// As(target) returns true. In the latter case, the As method is responsible for
// setting target.
//
// As will panic if target is not a non-nil pointer to either a type that implements
// error, or to any interface type. As returns false if err is nil.
//
// link to stdlib errors.As
func As(err error, target interface{}) bool

//go:linkname Unwrap errors.Unwrap

// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
func Unwrap(err error) error
