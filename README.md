[![Documentation](https://pkg.go.dev/badge/github.com/nikandfor/errors)](https://pkg.go.dev/github.com/nikandfor/errors?tab=doc)
[![Build Status](https://travis-ci.com/nikandfor/errors.svg?branch=master)](https://travis-ci.com/nikandfor/errors)
[![CircleCI](https://circleci.com/gh/nikandfor/errors.svg?style=svg)](https://circleci.com/gh/nikandfor/errors)
[![codecov](https://codecov.io/gh/nikandfor/errors/branch/master/graph/badge.svg)](https://codecov.io/gh/nikandfor/errors)
[![GolangCI](https://golangci.com/badges/github.com/nikandfor/errors.svg)](https://golangci.com/r/github.com/nikandfor/errors)
[![Go Report Card](https://goreportcard.com/badge/github.com/nikandfor/errors)](https://goreportcard.com/report/github.com/nikandfor/errors)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/nikandfor/errors?sort=semver)

# errors

Stdlib `errors` package extension. `go1.13` `errors.Is` and `errors.As` are the same functions as in stdlib (not even copies).

```go
// as usual
err = errors.New("msg")

// do not capture caller info
err = errors.NewNoLoc("msg")

// fmt.Sprintf like
err = errors.New("message %v", "args")

// one Frame higher
err = errors.NewDepth(1, "msg")

// the same result as previous
pc := errors.Caller(1)
err = errors.NewLoc(pc, "msg")

// Wrap error
err = errors.Wrap(err, "msg %v", "args")

// all the same function types are available
err = errors.WrapNoLoc(err, "msg")

err = errors.WrapDepth(err, 1, "msg %v", "args")

err = errors.WrapLoc(err, pc, "msg %v", "args")
```

## Caller

Caller frame can be added to error so later you can get to know where error was generated.

```go
f := errors.Caller(1)

f = errors.Funcentry(1)
```
