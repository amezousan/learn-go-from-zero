# Integer

## What to do?

Go through https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/integers

## Test Results

```
$ go test -v
=== RUN   TestFourOps
=== RUN   TestFourOps/Add:_Plus_+_Plus_become_plus
=== RUN   TestFourOps/Add:_Plus_=_Minus_become_zero
=== RUN   TestFourOps/Add:_Plus_<_Minus_become_minus
=== RUN   TestFourOps/Sub:_Plus_-_Minus_become_plus
=== RUN   TestFourOps/Multiply:_Minus_*_Minus_become_plus
=== RUN   TestFourOps/Divide:_Minus_/_Minus_become_plus
--- PASS: TestFourOps (0.00s)
    --- PASS: TestFourOps/Add:_Plus_+_Plus_become_plus (0.00s)
    --- PASS: TestFourOps/Add:_Plus_=_Minus_become_zero (0.00s)
    --- PASS: TestFourOps/Add:_Plus_<_Minus_become_minus (0.00s)
    --- PASS: TestFourOps/Sub:_Plus_-_Minus_become_plus (0.00s)
    --- PASS: TestFourOps/Multiply:_Minus_*_Minus_become_plus (0.00s)
    --- PASS: TestFourOps/Divide:_Minus_/_Minus_become_plus (0.00s)
PASS
ok  	integer	0.191s
```