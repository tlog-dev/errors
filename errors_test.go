package errors

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapperError(t *testing.T) {
	assert.EqualError(t, New("qwe"), "qwe")
	assert.EqualError(t, NewLoc("qwe %v %v", 1, 2), "qwe 1 2")
	assert.EqualError(t, Wrap(New("qwe %v %v", 1, 2), "context %v %v", "a", "b"), "context a b: qwe 1 2")
	assert.EqualError(t, WrapLoc(New("qwe %v %v", 1, 2), "context %v %v", "a", "b"), "context a b: qwe 1 2")
	assert.EqualError(t, Wrap(New("qwe %v %v", 1, 2), "context").(interface{ Unwrap() error }).Unwrap(), "qwe 1 2")
	assert.EqualError(t, Unwrap(Wrap(New("qwe %v %v", 1, 2), "context")), "qwe 1 2")
}

func TestIsUnwrap(t *testing.T) {
	mid := Wrap(os.ErrNotExist, "middle")
	err := Wrap(mid, "global")

	assert.True(t, Is(mid, os.ErrNotExist))
	assert.True(t, Is(err, os.ErrNotExist))

	e := Unwrap(mid)
	assert.True(t, e == os.ErrNotExist, e)

	e = Unwrap(err)
	assert.True(t, e == mid, e)
}
