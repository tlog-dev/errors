package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepthLoc(t *testing.T) {
	assert.Equal(t, "errors_loc_test.go:10", NewDepth(0, "msg").(wrapper).Location().String())
	assert.Equal(t, "errors_loc_test.go:11", NewLoc(Caller(0), "msg").(wrapper).Location().String())

	err := New("inner")

	assert.Equal(t, "errors_loc_test.go:15", WrapDepth(err, 0, "msg").(wrapper).Location().String())
	assert.Equal(t, "errors_loc_test.go:16", WrapLoc(err, Caller(0), "msg").(wrapper).Location().String())
}
