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
go test -race         # Check RACE condition
go test -bench=.
godoc -http=:8080     # http://localhost:8080/pkg/<...>
```

***

## Useful Articles
* [Go](https://wiki.nikiv.dev/programming-languages/go/)
* [--> Go Recipes - Amazing tools](https://github.com/nikolaydubina/go-recipes)
* [--> The Busy Developer's Guide to Go Profiling, Tracing and Observability](https://github.com/DataDog/go-profiler-notes/blob/main/guide/README.md)
* [Constant errors](https://dave.cheney.net/2016/04/07/constant-errors)
* [Comprehensive Guide to Testing in Go](https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/)
* [Testing Go code](https://subscription.packtpub.com/book/programming/9781838559335/11/ch11lvl1sec145/testing-go-code)
* [Golang basics - writing unit tests](https://blog.alexellis.io/golang-writing-unit-tests/)
* [Testing Your Go App: Get Started The Right Way](https://www.toptal.com/go/your-introductory-course-to-testing-with-go)
* [Better tests for Golang apps](https://levelup.gitconnected.com/better-tests-for-golang-apps-681ed2338677)

* [Go Blog - The Laws of Reflection](https://go.dev/blog/laws-of-reflection)
* [Go Blog - Context](https://go.dev/blog/context)

***

## Go Profiling
* [How to profile Go with pprof in 30 seconds](https://dev.to/agamm/how-to-profile-go-with-pprof-in-30-seconds-592a)
* [Getting started with Go CPU and memory profiling](https://flaviocopes.com/golang-profiling/)
* [Profiling in GoLang](https://golangdocs.com/profiling-in-golang)
* [Make Your Go Programs Lightning Fast With Profiling](https://code.tutsplus.com/tutorials/make-your-go-programs-lightning-fast-with-profiling--cms-29809)
* [Profiling in Go](https://betterprogramming.pub/profiling-in-go-78cf71f81a07)
* [pprof++: A Go Profiler with Hardware Performance Monitoring](https://www.uber.com/blog/pprof-go-profiler/)
* [Performance Measuring, Profiling, and Optimizing Tips for Go Web Applications](https://articles.wesionary.team/performance-measuring-profiling-and-optimizing-tips-for-go-web-applications-20f2f812ff6e)
