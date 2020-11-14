package fib

func Fib1(n int) int {
	if n < 2 {
		return n
	}
	return Fib1(n-1) + Fib1(n-2)
}

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
