package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMask(t *testing.T) {
	a := New("a")
	b := New("mask")

	m := Mask(a, b)

	assert.Equal(t, b.Error(), m.Error())

	assert.True(t, Is(m, b))
	assert.False(t, Is(m, a))

	assert.True(t, Unwrap(m) == b)

	um, ok := m.(interface{ Unmask() error })
	assert.True(t, ok)
	assert.True(t, um.Unmask() == a)
}
