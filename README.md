# Common util or extension  [![Build Status](https://travis-ci.org/go-eden/common.svg?branch=master)](https://travis-ci.org/go-eden/common)

This library provides some utils or extension.

# Install

```shell
go get github.com/go-eden/common
```

# Modules

## [`etime`](./etime)

`etime` extend `golang`'s `time` package, and it should be faster than `time`. In my benchmark, the performance
of `etime.Now*` was about `40 ns/op`, `time.Now()` was about `68 ns/op`:

## [`emath`](./emath)

some math utils.

## [`esync`](./esync)

Provides an implementation of `reentrant lock` or `reentrant mutex`.

Wraps some atomic number:

+ `AtomicBool`
+ `AtomicInt`
+ `AtomicUint16`
+ `AtomicFloat32`
+ `AtomicFloat64`

## [`goid`](./goid)

`goid` provides `Gid()` function to find `ID` of the current `goroutine`.

# License

MIT