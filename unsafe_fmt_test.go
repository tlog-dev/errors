package errors

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	testformatter struct {
		flags [64]bool

		wid int
		prc int

		widok bool
		prcok bool

		verb rune

		buf []byte
	}
)

func TestSubPrintArg(t *testing.T) {
	var b bytes.Buffer

	var f testformatter

	fmt.Fprintf(&b, "%+012.6q", &f)

	assert.Equal(t, testformatter{
		flags: flags("+0"),
		wid:   12,
		widok: true,
		prc:   6,
		prcok: true,
		verb:  'q',
	}, f)

	//

	var f2 testformatter

	subPrintArg(&f, &f2, 'v')

	assert.Equal(t, testformatter{
		flags: flags("+0"),
		wid:   12,
		widok: true,
		prc:   6,
		prcok: true,
		verb:  'v',
	}, f2)

	//

	f = testformatter{}
	f2 = testformatter{
		flags: flags("-# "),
		prc:   3,
		prcok: true,
	}

	subPrintArg(&f2, &f, 'v')

	assert.Equal(t, testformatter{
		flags: flags(" -#"),
		prc:   3,
		prcok: true,
		verb:  'v',
	}, f)
}

func BenchmarkPringArg(b *testing.B) {
	b.ReportAllocs()

	f := testformatter{}

	for i := 0; i < b.N; i++ {
		pp := newPrinter()
		printArg(pp, &f, 'v')
		ppFree(pp)
	}
}

func BenchmarkPringArgFallback(b *testing.B) {
	b.ReportAllocs()

	f := testformatter{}
	f2 := testformatter{
		flags: flags("-# "),
		prc:   3,
		prcok: true,
	}

	for i := 0; i < b.N; i++ {
		subPrintArg(&f2, &f, 'v')
	}
}

func (f *testformatter) Format(s fmt.State, verb rune) {
	f.flags = [64]bool{}

	for _, q := range "-+# 0" {
		if s.Flag(int(q)) {
			f.flags[q] = true
		}
	}

	f.wid, f.widok = s.Width()
	f.prc, f.prcok = s.Precision()

	f.verb = verb
}

func (f *testformatter) Flag(c int) bool {
	return f.flags[c]
}

func (f *testformatter) Width() (int, bool) {
	return f.wid, f.widok
}

func (f *testformatter) Precision() (int, bool) {
	return f.prc, f.prcok
}

func (f *testformatter) Write(p []byte) (int, error) {
	f.buf = append(f.buf, p...)

	return len(p), nil
}

func flags(f string) (r [64]bool) {
	for _, q := range f {
		r[q] = true
	}

	return
}
