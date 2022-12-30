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
* The benchmark function takes one argument only `b *testing.B`

* In order to use the `*testing.T` type, you need to `import "testing"`, like we did with `fmt` in the other file

***
## Useful Articles

