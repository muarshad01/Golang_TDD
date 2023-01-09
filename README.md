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
* [Profiling Go programs with pprof](https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/)
* [How to profile Go with pprof in 30 seconds](https://dev.to/agamm/how-to-profile-go-with-pprof-in-30-seconds-592a)
* [Profiling in GoLang](https://golangdocs.com/profiling-in-golang)
* [Make Your Go Programs Lightning Fast With Profiling](https://code.tutsplus.com/tutorials/make-your-go-programs-lightning-fast-with-profiling--cms-29809)
* [Profiling in Go](https://betterprogramming.pub/profiling-in-go-78cf71f81a07)
* [pprof++: A Go Profiler with Hardware Performance Monitoring](https://www.uber.com/blog/pprof-go-profiler/)
* [Performance Measuring, Profiling, and Optimizing Tips for Go Web Applications](https://articles.wesionary.team/performance-measuring-profiling-and-optimizing-tips-for-go-web-applications-20f2f812ff6e)

***

```go
$ brew
$ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
$ brew install graphviz
```

```go
$ mkdir install_pprof
$ cd install_pprof
$ got mod init install_pprof
$ go get -u github.com/google/pprof
```

```go
$ go tool pprof http://localhost:{APP PORT}/{pprof endpoint}
$ go tool pprof http://localhost:6060/debug/pprof/heap
$ go tool pprof http://localhost:6060/debug/pprof/goroutine
$ go tool pprof http://localhost:6060/debug/pprof/block
$ go tool pprof http://localhost:6060/debug/pprof/mutex
$ go tool pprof http://localhost:6060/debug/pprof/profile
$ go tool pprof http://localhost:6060/debug/pprof/trace
```

* `/debug/pprof/alloc`:         A sampling of all past memory allocations
* `/debug/pprof/block`:         Stack traces that led to blocking on synchronization primitives
* `/debug/pprof/cmdline`:       The command line invocation of the current program
* `/debug/pprof/goroutine`:     Stack traces of all current goroutines
* `/debug/pprof/heap`:          A sampling of memory allocations of live objects. You can specify the gc GET parameter to run GC before taking the heap sample.
* `/debug/pprof/mutex`:         Stack traces of holders of contended mutexes
* `/debug/pprof/profile`:       CPU profile. You can specify the duration in the seconds GET parameter. After you get the profile file, use the go tool pprof command to investigate the profile.
* `/debug/pprof/threadcreate`:  Stack traces that led to the creation of new OS threads
* `/debug/pprof/trace`:         A trace of execution of the current program. You can specify the duration in the seconds GET parameter. After you get the trace file, use the go tool trace command to investigate the trace.
