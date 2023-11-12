package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"tlog.app/go/loc"
)

func TestDepthCaller(t *testing.T) {
	assert.Equal(t, "errors_location_test.go:11", NewDepth(0, "msg").(Callerer).Caller().String())              //nolint:errorlint
	assert.Equal(t, "errors_location_test.go:12", NewCaller(loc.Caller(0), "msg").(Callerer).Caller().String()) //nolint:errorlint

	err := New("inner")

	assert.Equal(t, "errors_location_test.go:16", WrapDepth(err, 0, "msg").(Callerer).Caller().String())              //nolint:errorlint
	assert.Equal(t, "errors_location_test.go:17", WrapCaller(err, loc.Caller(0), "msg").(Callerer).Caller().String()) //nolint:errorlint
}

func TestStackCallers(t *testing.T) {
	testStackCallers(t)
}

func testStackCallers(t *testing.T) {
	assert.Equal(t, "errors_location_test.go:25", NewStack(0, 2, "msg").(Callerer).Caller().String())                //nolint:errorlint
	assert.Equal(t, "errors_location_test.go:26", NewCallers(loc.Callers(0, 2), "msg").(Callerer).Caller().String()) //nolint:errorlint

	err := New("inner")

	assert.Equal(t, "errors_location_test.go:30", WrapStack(err, 0, 2, "msg").(Callerer).Caller().String())                //nolint:errorlint
	assert.Equal(t, "errors_location_test.go:31", WrapCallers(err, loc.Callers(0, 2), "msg").(Callerer).Caller().String()) //nolint:errorlint

	assert.Equal(t, PC(0), NewCallers(PCs{}, "msg").(Callerer).Caller())                                //nolint:errorlint
	assert.Equal(t, "errors_location_test.go:34", NewStack(0, 1, "msg").(Callerer).Caller().String())   //nolint:errorlint
	assert.Equal(t, "errors_location_test.go:35", NewStack(0, 1, "msg").(Callerser).Callers().String()) //nolint:errorlint
}
