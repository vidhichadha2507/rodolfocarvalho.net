---
title: "How to Pipe Go Benchmarks to Benchstat"
slug: "go-test-bench-pipe-to-benchstat"
description: ""
categories: []
tags: [golang, benchmarking, shell, testing]
series: []
date: 2020-11-30T18:33:38+01:00
---

A quick note to remind myself how to pipe the output of `go test -bench` to
standard output and `benchstat`.

<!--more-->

The trick is to use `/dev/stdin` as the input file for `benchstat`, as
[`benchstat` doesn't read from standard
input](https://go-review.googlesource.com/c/perf/+/94905/).

```shell
go test -run='^$' -bench=. -benchmem -count=5 | tee >(benchstat /dev/stdin)
```

Example run:

```console
$ go test -run='^$' -bench=SetTag -benchmem -count=5 | tee >(benchstat /dev/stdin)
goos: darwin
goarch: amd64
pkg: github.com/getsentry/sentry-go
BenchmarkSetTag/SetTag-16               123519501                9.92 ns/op            0 B/op          0 allocs/op
BenchmarkSetTag/SetTag-16               121863286                9.63 ns/op            0 B/op          0 allocs/op
BenchmarkSetTag/SetTag-16               124801513               10.2 ns/op             0 B/op          0 allocs/op
BenchmarkSetTag/SetTag-16               119316870                9.97 ns/op            0 B/op          0 allocs/op
BenchmarkSetTag/SetTag-16               123358520                9.69 ns/op            0 B/op          0 allocs/op
BenchmarkSetTag/SetTag2-16              91541181                12.0 ns/op             0 B/op          0 allocs/op
BenchmarkSetTag/SetTag2-16              86802069                12.0 ns/op             0 B/op          0 allocs/op
BenchmarkSetTag/SetTag2-16              89232990                12.0 ns/op             0 B/op          0 allocs/op
BenchmarkSetTag/SetTag2-16              100462875               11.9 ns/op             0 B/op          0 allocs/op
BenchmarkSetTag/SetTag2-16              93366249                12.8 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/getsentry/sentry-go  17.740s
name               time/op
SetTag/SetTag-16   9.88ns ± 3%
SetTag/SetTag2-16  12.0ns ± 1%

name               alloc/op
SetTag/SetTag-16    0.00B
SetTag/SetTag2-16   0.00B

name               allocs/op
SetTag/SetTag-16     0.00
SetTag/SetTag2-16    0.00
```
