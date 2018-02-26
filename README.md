# builderpool
[![Build Status](https://travis-ci.org/nasa9084/go-builderpool.svg?branch=master)](https://travis-ci.org/nasa9084/go-builderpool)
[![Godoc](https://godoc.org/github.com/nasa9084/go-builderpool?status.svg)](https://godoc.org/github.com/nasa9084/go-builderpool)
---

A simple strings.Builder pool using sync.Pool inspired by [lestrrat-go/bufferpool](https://github.com/lestrrat-go/bufferpool).

## SYNOPSIS

``` go
import "github.com/nasa9084/go-builderpool"

var pool = builderpool.New()

func main() {
    builder := pool.Get()
    defer pool.Release(builder)

    // ...
}
```
