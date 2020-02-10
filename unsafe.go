package errors

import (
	_ "errors"
	_ "unsafe"
)

//go:linkname Is errors.Is
func Is(err, target error)

//go:linkname As errors.Is
func As(err, target error)
