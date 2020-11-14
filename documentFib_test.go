//Package provides implementation
package fib_test

import (
	"fmt"

	"github.com/RomanAvdeenko/fib"
)

func ExampleFib() {
	fmt.Println(fib.Fib1(3))
	fmt.Println(fib.Fib(8))
	//Output:
	//2
	//21
}
