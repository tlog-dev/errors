[![Documentation](https://pkg.go.dev/badge/tlog.app/go/errors)](https://pkg.go.dev/tlog.app/go/errors?tab=doc)
[![Build Status](https://travis-ci.com/tlog-dev/errors.svg?branch=master)](https://travis-ci.com/tlog-dev/errors)
[![CircleCI](https://circleci.com/gh/tlog-dev/errors.svg?style=svg)](https://circleci.com/gh/tlog-dev/errors)
[![codecov](https://codecov.io/gh/tlog-dev/errors/branch/master/graph/badge.svg)](https://codecov.io/gh/tlog-dev/errors)
[![Go Report Card](https://goreportcard.com/badge/tlog.app/go/errors)](https://goreportcard.com/report/tlog.app/go/errors)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/tlog-dev/errors?sort=semver)

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
pc := loc.Caller(1)
err = errors.NewLoc(pc, "msg")

// Wrap error
err = errors.Wrap(err, "msg %v", "args")

// all the same function types are available
err = errors.WrapNoLoc(err, "msg")

err = errors.WrapDepth(err, 1, "msg %v", "args")

err = errors.WrapLoc(err, pc, "msg %v", "args")
```

## Caller

Caller frame can be added to error so later you can get to know where error was generated. It's added by default and captures instruction calling `errors.(Wrap|New)*`.

Caller is moved to a separate module [github.com/nikandfor/loc](https://github.com/nikandfor/loc).

```go
pc := loc.Caller(1)

pc = loc.FuncEntry(1)
```
