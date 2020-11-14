package fib_test

import (
	"testing"

	"github.com/RomanAvdeenko/fib"
)

var result int

func benchmarkFib(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fib.Fib(n)
	}
	result = r
}

func benchmarkFib1(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fib.Fib1(n)
	}
	result = r
}

func Benchmark10Fib(b *testing.B) {
	benchmarkFib(b, 10)
}

func Benchmark10Fib1(b *testing.B) {
	benchmarkFib1(b, 10)
}

func Benchmark30Fib(b *testing.B) {
	benchmarkFib(b, 30)
}

func Benchmark30Fib1(b *testing.B) {
	benchmarkFib1(b, 30)
}

func Benchmark50Fib(b *testing.B) {
	benchmarkFib(b, 50)
}

func Benchmark50Fib1(b *testing.B) {
	benchmarkFib1(b, 50)
}
