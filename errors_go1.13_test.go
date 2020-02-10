// +build go1.13

package errors

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIs(t *testing.T) {
	mid := Wrap(os.ErrNotExist, "middle")
	err := Wrap(mid, "global")

	assert.True(t, Is(mid, os.ErrNotExist))
	assert.True(t, Is(err, os.ErrNotExist))
}
