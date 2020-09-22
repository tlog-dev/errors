package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepthFrame(t *testing.T) {
	assert.Equal(t, "errors_frame_test.go:10", NewDepth(0, "msg").(wrapper).Frame().String())
	assert.Equal(t, "errors_frame_test.go:11", NewFrame(Caller(0), "msg").(wrapper).Frame().String())

	err := New("inner")

	assert.Equal(t, "errors_frame_test.go:15", WrapDepth(err, 0, "msg").(wrapper).Frame().String())
	assert.Equal(t, "errors_frame_test.go:16", WrapFrame(err, Caller(0), "msg").(wrapper).Frame().String())
}
