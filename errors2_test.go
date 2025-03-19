package errors_test

import (
	stderrors "errors"
	"fmt"
	"testing"

	"tlog.app/go/errors"
)

type testErr struct {
	val int
}

func TestErrors(tb *testing.T) {
	err1 := errors.New("some error")
	if err1 == nil {
		tb.Errorf("error expected")
	}

	err1dup := errors.New("some error")
	if err1dup == nil {
		tb.Errorf("error expected")
	}
	if err1dup == err1 { //nolint:errorlint
		tb.Errorf("expected to be a separate error")
	}

	if msg := err1.Error(); msg != "some error" {
		tb.Errorf("unexpected message: %v", msg)
	}

	err2 := errors.New("some error: %v", 5)
	if err2 == nil {
		tb.Errorf("error expected")
	}

	err1w := errors.Wrap(err1, "wrap1")
	if err1w == nil {
		tb.Errorf("error wrapper expected")
	}

	if msg := err1w.Error(); msg != "wrap1: some error" {
		tb.Errorf("unexpected message: %v", msg)
	}

	if unerr := errors.Unwrap(err1w); unerr != err1 { //nolint:errorlint
		tb.Errorf("err1w unwrap expected to be err1")
	}

	if !errors.Is(err1w, err1) {
		tb.Errorf("expected err1w to be err1 wrapper")
	}
	if !stderrors.Is(err1w, err1) {
		tb.Errorf("expected err1w to be err1 wrapper (stdlib)")
	}

	if errors.Is(err1w, err2) {
		tb.Errorf("err1w is not err2 wrapper")
	}
	if errors.Is(err2, err1) {
		tb.Errorf("err2 is not err1 wrapper")
	}

	err3 := &testErr{val: 3}
	err3w := errors.Wrap(err3, "wrapped")
	var err3target *testErr

	if ok := errors.As(err3w, &err3target); !ok || err3 != err3target {
		tb.Errorf("unwrapped err3w expected to be err3, got ok %v  err %v  equal %v", ok, err3target, err3target == err3)
	}

	func() {
		defer func() {
			p := recover()
			if p == nil {
				tb.Errorf("panic expected")
			}
		}()

		_ = errors.Wrap(nil, "wrap nil")
	}()
}

func (e *testErr) Error() string { return fmt.Sprintf("testErr(%v)", e.val) }
