[![Documentation](https://pkg.go.dev/badge/tlog.app/go/errors)](https://pkg.go.dev/tlog.app/go/errors?tab=doc)
[![Go workflow](https://github.com/tlog-dev/errors/actions/workflows/go.yml/badge.svg)](https://github.com/tlog-dev/errors/actions/workflows/go.yml)
[![CircleCI](https://circleci.com/gh/tlog-dev/errors.svg?style=svg)](https://circleci.com/gh/tlog-dev/errors)
[![codecov](https://codecov.io/gh/tlog-dev/errors/branch/master/graph/badge.svg)](https://codecov.io/gh/tlog-dev/errors)
[![Go Report Card](https://goreportcard.com/badge/tlog.app/go/errors)](https://goreportcard.com/report/tlog.app/go/errors)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/tlog-dev/errors?sort=semver)

# errors

`errors` is a wrapper around the standard `errors.New` and `fmt.Errorf` functions.
It unifies their APIs, allowing you to create new errors or wrap existing ones with rich context (formatted message with arguments).

It provides two core functions:

```
errors.New(format string, args ...any) error
errors.Wrap(err error, format string, args ...any) error
```

While it previously had more features, I eventually realized that,
in most cases, anything beyond simple error wrapping is an overcomplication.
