## learn-go-with-tests

* [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)

***

## Writing tests

* Writing a test is just like writing a function, with a few rules
* It needs to be in a file with a name like `xxx_test.go`
* The test function must start with the word `Test`
  - The test function takes one argument only `t *testing.T`
* The example function must start with the word `Example`
* The benchmark function must start with the word `Benchmark`
  - The benchmark function takes one argument only `b *testing.B`
  - The `testing.B` gives you access to the cryptically named `b.N`.

* In order to use the `*testing.T` type, you need to `import "testing"`, like we did with `fmt` in the other file

***

## errcheck

* [errcheck](https://github.com/kisielk/errcheck)

```go
$ mkdir -p $GOPATH/src/<folder>             # E.g., install_errcheck
$ cd $GOPATH/src/<folder>
$ go mod init <folder>
$ go mod tidy
$ go install github.com/kisielk/errcheck@latest

$ errcheck .
$ errcheck ./...
$ errcheck all
```
***

## run.sh

```go
#!/bin/bash

go fmt ./...
go vet ./...
errcheck ./...

go test -v
go test -bench=.
godoc -http=:8080     # http://localhost:8080/pkg/<...>
```

***

## Useful Articles

* [Constant errors](https://dave.cheney.net/2016/04/07/constant-errors)
* [Comprehensive Guide to Testing in Go](https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/)

***
