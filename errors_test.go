package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	assert.EqualError(t, New("qwe"), "qwe")
	assert.EqualError(t, Newf("qwe"), "qwe")
	assert.EqualError(t, Newf("qwe %v %v", 1, 2), "qwe 1 2")
	assert.EqualError(t, Wrap(Newf("qwe %v %v", 1, 2), "context"), "context: qwe 1 2")
	assert.EqualError(t, Wrapf(Newf("qwe %v %v", 1, 2), "context %v %v", "a", "b"), "context a b: qwe 1 2")
	assert.EqualError(t, WrapfNoLoc(Newf("qwe %v %v", 1, 2), "context %v %v", "a", "b"), "context a b: qwe 1 2")
	assert.EqualError(t, Cause(Wrap(Newf("qwe %v %v", 1, 2), "context")), "qwe 1 2")
}
