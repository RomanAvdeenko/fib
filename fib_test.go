package fib_test

import (
	"testing"
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal()
		})
	}

}
