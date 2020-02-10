package errors

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapperError(t *testing.T) {
	assert.EqualError(t, New("qwe"), "qwe")
	assert.EqualError(t, NewHere("qwe %v %v", 1, 2), "qwe 1 2")
	assert.EqualError(t, Wrap(New("qwe %v %v", 1, 2), "context %v %v", "a", "b"), "context a b: qwe 1 2")
	assert.EqualError(t, WrapHere(New("qwe %v %v", 1, 2), "context %v %v", "a", "b"), "context a b: qwe 1 2")
	assert.EqualError(t, Wrap(New("qwe %v %v", 1, 2), "context").(interface{ Unwrap() error }).Unwrap(), "qwe 1 2")
	assert.EqualError(t, Unwrap(Wrap(New("qwe %v %v", 1, 2), "context")), "qwe 1 2")
}

func TestUnwrap(t *testing.T) {
	mid := Wrap(os.ErrNotExist, "middle")
	err := Wrap(mid, "global")

	e := Unwrap(mid)
	assert.True(t, e == os.ErrNotExist, e)

	e = Unwrap(err)
	assert.True(t, e == mid, e)
}
