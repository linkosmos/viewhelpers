# viewhelpers
HTML template helpers

[![Build Status](https://travis-ci.org/linkosmos/viewhelpers.svg?branch=master)](https://travis-ci.org/linkosmos/viewhelpers)
[![Coverage Status](https://coveralls.io/repos/github/linkosmos/viewhelpers/badge.svg?branch=master)](https://coveralls.io/github/linkosmos/viewhelpers?branch=master)
[![GoDoc](http://godoc.org/github.com/linkosmos/viewhelpers?status.svg)](http://godoc.org/github.com/linkosmos/viewhelpers)
[![Go Report Card](http://goreportcard.com/badge/linkosmos/viewhelpers)](http://goreportcard.com/report/linkosmos/viewhelpers)
[![BSD License](http://img.shields.io/badge/license-BSD-blue.svg)](http://opensource.org/licenses/BSD-3-Clause)


### Usage

```go
var App template.FuncMap = template.FuncMap{
  "loadtime":                  viewhelpers.Loadtime,
  "dict":                      viewhelpers.Dict,
  "len":                       viewhelpers.Len,
}

```
