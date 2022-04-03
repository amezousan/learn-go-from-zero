# HelloWorld

## What to do?
1. Refer to https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world
2. Make a test case

## How to use?
```
$ go run hello
Hello, world

# If there is no go.sum, you need to run go mod init SOMENAME
# in each new folder before running commands like go test or go build.
$ go mod init hello
go: creating new go.mod: module hello
go: to add module requirements and sums:
	go mod tidy

$ go test -v
=== RUN   TestHello
=== RUN   TestHello/saying_hello_to_people
=== RUN   TestHello/say_'Hello,_World'_when_an_empty_string_is_supplied
--- PASS: TestHello (0.00s)
    --- PASS: TestHello/saying_hello_to_people (0.00s)
    --- PASS: TestHello/say_'Hello,_World'_when_an_empty_string_is_supplied (0.00s)
PASS
ok  	hello	0.159s
```

### Ref
* [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world)
* [Goのtestingパッケージの基本を理解する](https://qiita.com/taisa831/items/85fea8d970bcadd796b9)