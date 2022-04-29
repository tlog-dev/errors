package errors

type (
	masked struct {
		err  error
		mask error
	}
)

// Mask masks original error err with error mask.
// Masked errors is almost identical to mask but allows to recover original error.
// Unmasking is similar to Unwrapping and is done by method
//
//     Unmask() error
//
func Mask(err, mask error) error {
	return masked{
		err:  err,
		mask: mask,
	}
}

func (e masked) Error() string { return e.mask.Error() }

func (e masked) Unwrap() error { return e.mask }

func (e masked) Unmask() error { return e.err }
