package fib_test

import (
	"testing"

	"github.com/RomanAvdeenko/fib"
	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	testCases := []struct {
		name string
		in   int
		want int
	}{
		{
			name: "zero",
			in:   0,
			want: 0,
		}, {
			name: "one",
			in:   1,
			want: 1,
		},
		{
			name: "two",
			in:   2,
			want: 1,
		},
		{
			name: "three",
			in:   3,
			want: 2,
		},
		{
			name: "four",
			in:   4,
			want: 3,
		}, {
			name: "five",
			in:   5,
			want: 5,
		},
		{
			name: "six",
			in:   6,
			want: 8,
		}, {
			name: "nine",
			in:   9,
			want: 34,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, fib.Fib(tc.in))
		})
	}

}
