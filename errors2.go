// Package errors is a minimal wrapper around the standard library errors package.
// It changes the API from New(msg string) and fmt.Errorf(format string, args ...any)
// to unified New(format string, args ...any) and Wrap(err error, format string, args ...any) functions.
package errors

import (
	"errors"
	"fmt"
)

type (
	wrap struct {
		err error
		msg string
	}
)

func Is(err, target error) bool     { return errors.Is(err, target) }
func As(err error, target any) bool { return errors.As(err, target) }
func Unwrap(err error) error        { return errors.Unwrap(err) }

func New(format string, args ...any) error {
	if args == nil {
		return errors.New(format)
	}

	return fmt.Errorf(format, args...)
}

func Wrap(err error, format string, args ...any) error {
	if err == nil {
		panic("wrapping nil error")
	}

	return &wrap{
		err: err,
		msg: fmt.Sprintf(format, args...),
	}
}

func (w *wrap) Error() string {
	return fmt.Sprintf("%v: %v", w.msg, w.err)
}

func (w *wrap) Unwrap() error {
	return w.err
}
