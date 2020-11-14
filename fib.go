// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fib implements functions for the manipulation of byte slices.
// It is analogous to the facilities of the strings package.
package fib

// Fib1 is a canonical slow fibonacci implementation.
func Fib1(n int) int {
	if n < 2 {
		return n
	}
	return Fib1(n-1) + Fib1(n-2)
}

// Fib is a  caching map fibonacci implementation.
func Fib(n int) int {
	fn := make(map[int]int, n+1)
	for i := 0; i <= n; i++ {
		if i < 2 {
			fn[i] = i
		} else {
			fn[i] = fn[i-1] + fn[i-2]
		}
	}
	return fn[n]
}
