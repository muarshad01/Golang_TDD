## learn-go-with-tests

* [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)

***

## Writing tests

* Writing a test is just like writing a function, with a few rules
* It needs to be in a file with a name like `xxx_test.go`
* The test function must start with the word `TestYYY(t *testing.T)`
  - The test function takes one argument only `t *testing.T`
* The example function must start with the word `ExampleYYY()`
* The benchmark function must start with the word `BenchmarkYYY(b *testing.B)`
  - The benchmark function takes one argument only `b *testing.B`
  - The `testing.B` gives you access to the cryptically named `b.N`.

* In order to use the `*testing.T` type, you need to `import "testing"`, like we did with `fmt` in the other file

***
## Useful Articles

