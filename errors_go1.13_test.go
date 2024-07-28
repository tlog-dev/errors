//go:build go1.13
// +build go1.13

package errors

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	asError string
)

func TestIs(t *testing.T) {
	mid := Wrap(os.ErrNotExist, "middle")
	err := Wrap(mid, "global")

	assert.True(t, Is(mid, os.ErrNotExist))
	assert.True(t, Is(err, os.ErrNotExist))
}

func TestAs(t *testing.T) {
	err := asError("as_error")
	mid := Wrap(err, "middle")
	top := Wrap(mid, "top")

	assert.True(t, Is(top, asError("as_error")))

	var target asError

	assert.True(t, As(top, &target))
	assert.Equal(t, err, target)
}

func (e asError) Error() string { return string(e) }
