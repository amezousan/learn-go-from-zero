# Ref

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration

# Test Results

```
$ go test -v
=== RUN   TestRepeat
=== RUN   TestRepeat/a_x_5
=== RUN   TestRepeat/0_x_5
--- PASS: TestRepeat (0.00s)
    --- PASS: TestRepeat/a_x_5 (0.00s)
    --- PASS: TestRepeat/0_x_5 (0.00s)
PASS
ok  	iteration	0.136s

$ go test -bench=.
goos: darwin
goarch: arm64
pkg: iteration
BenchmarkRepeat-8   	13118445	        84.52 ns/op
PASS
ok  	iteration	1.297s
```