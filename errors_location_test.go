package errors

import (
	"testing"

	"github.com/nikandfor/loc"
	"github.com/stretchr/testify/assert"
)

func TestDepthLocation(t *testing.T) {
	assert.Equal(t, "errors_location_test.go:11", NewDepth(0, "msg").(Locationer).Location().String())
	assert.Equal(t, "errors_location_test.go:12", NewLoc(loc.Caller(0), "msg").(Locationer).Location().String())

	err := New("inner")

	assert.Equal(t, "errors_location_test.go:16", WrapDepth(err, 0, "msg").(Locationer).Location().String())
	assert.Equal(t, "errors_location_test.go:17", WrapLoc(err, loc.Caller(0), "msg").(Locationer).Location().String())
}
