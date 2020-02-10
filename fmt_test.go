package errors

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorFormat(t *testing.T) {
	err := NewLoc("inner")

	assert.Equal(t, "inner (fmt_test.go:12)", fmt.Sprintf("%+v", err))

	// more
	err = New("inner")

	assert.Equal(t, "inner", fmt.Sprintf("%v", err))
	assert.Equal(t, "inner", fmt.Sprintf("%+v", err))
	assert.Equal(t, "inner", fmt.Sprintf("% +v", err))

	// more
	err = WrapLoc(WrapLoc(NewLoc("inner"), "middle"), "global")

	assert.Equal(t, "global: middle: inner", err.Error())

	assert.Equal(t, err.Error(), fmt.Sprintf("%v", err))

	assert.Equal(t, "global (fmt_test.go:24): middle (fmt_test.go:24): inner (fmt_test.go:24)", fmt.Sprintf("%+v", err))

	assert.Equal(t, `global
at github.com/nikandfor/errors/fmt_test.go:24
middle
at github.com/nikandfor/errors/fmt_test.go:24
inner
at github.com/nikandfor/errors/fmt_test.go:24`, fmt.Sprintf("% +v", err))

	// one more
	err = Wrap(Wrap(os.ErrNotExist, "middle"), "global")

	assert.Equal(t, "global: middle: file does not exist", err.Error())

	assert.Equal(t, err.Error(), fmt.Sprintf("%v", err))

	assert.Equal(t, "global: middle: file does not exist", fmt.Sprintf("%+v", err))

	assert.Equal(t, "global: middle: file does not exist", fmt.Sprintf("% +v", err))

	// one more
	err = Wrap(Wrap(NewLoc("inner"), "middle"), "global")

	assert.Equal(t, "global: middle: inner", err.Error())

	assert.Equal(t, err.Error(), fmt.Sprintf("%v", err))

	assert.Equal(t, "global: middle: inner (fmt_test.go:51)", fmt.Sprintf("%+v", err))

	assert.Equal(t, "global: middle: inner\nat github.com/nikandfor/errors/fmt_test.go:51", fmt.Sprintf("% +v", err))
}
