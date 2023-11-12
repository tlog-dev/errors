//nolint:errorlint
package errors

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"tlog.app/go/loc"
)

func TestError(t *testing.T) {
	assert.Equal(t, "qwe", New("qwe").Error())
	assert.Equal(t, nomessage, New("").Error())
	assert.Equal(t, "qwe", Wrap(New("qwe"), "").Error())
}

type testWrapper struct { //nolint:errname
	err error
}

func TestWrapperError(t *testing.T) {
	assert.EqualError(t, New("qwe"), "qwe")
	assert.EqualError(t, NewNoCaller("qwe %v %v", 1, 2), "qwe 1 2")
	assert.EqualError(t, Wrap(New("qwe %v %v", 1, 2), "context %v %v", "a", "b"), "context a b: qwe 1 2")
	assert.EqualError(t, WrapNoCaller(New("qwe %v %v", 1, 2), "context %v %v", "a", "b"), "context a b: qwe 1 2")
	assert.EqualError(t, Wrap(New("qwe %v %v", 1, 2), "context").(interface{ Unwrap() error }).Unwrap(), "qwe 1 2")
	assert.EqualError(t, Unwrap(Wrap(New("qwe %v %v", 1, 2), "context")), "qwe 1 2")
}

//nolint:goerr113
func TestUnwrap(t *testing.T) {
	mid := Wrap(os.ErrNotExist, "middle")
	err := Wrap(mid, "global")

	assert.True(t, os.ErrNotExist == Unwrap(mid))

	assert.True(t, mid == Unwrap(err))

	assert.True(t, mid == Unwrap(testWrapper{mid}))

	assert.NoError(t, Unwrap(os.ErrNotExist))
}

func TestWrapNil(t *testing.T) {
	var err error

	assert.NoError(t, Wrap(err, "qwe"))
	assert.NoError(t, WrapNoCaller(err, "qwe"))
	assert.NoError(t, WrapDepth(err, 0, "qwe"))
	assert.NoError(t, WrapStack(err, 0, 0, "qwe"))
	assert.NoError(t, WrapCaller(err, loc.FuncEntry(0), "qwe"))
	assert.NoError(t, WrapCallers(err, loc.Callers(0, 1), "qwe"))
}

func (w testWrapper) Error() string { return "none" }

func (w testWrapper) Unwrap() error { return w.err }
