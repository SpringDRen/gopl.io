//Exercise 1.3: Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses strings.Join.
//
// Run with go test -bench=.
//The "." pattern causes it to match all benchmarks in the word package
package concat_test

//*_test.go files are a part of it when built by go test.

//bench test resut:
//goos: windows
//goarch: amd64
//pkg: github.com/SpringDRen/gopl.io/ch01/1.2/ex1.3
//BenchmarkConcat-4        2000000               758 ns/op
//BenchmarkStrJoin-4       5000000               240 ns/op
//PASS
//ok      github.com/SpringDRen/gopl.io/ch01/1.2/ex1.3    4.085s

import (
	"testing"
	"strings"
)

var args = []string{"911", "what's", "your", "emergency", "Can", "you", "hear", "me"}

func concat(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
}

//In Go, a benchmark function looks like a test function, but with the Benchmark prefix and a *testing.B parameter

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat(args)
	}
}

func BenchmarkStrJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(args, " ")
	}
}
