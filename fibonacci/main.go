package main

import "fmt"

func main() {
	fmt.Println(fib2(10))
	fmt.Println(fib3(40))
	fmt.Println(fib4(40))
}

// fib1 is a slow recursive implementation of the Fibonacci sequence
// it never finishes
func fib1(n int) int {
	return fib1(n-1) + fib1(n-2)
}

func fib2(n int) int {
	if n < 2 {
		return n
	}
	return fib2(n-1) + fib2(n-2)
}

var memo map[int64]int64

func fib3(n int64) int64 {
	if memo == nil {
		memo = make(map[int64]int64)
		memo[0] = 0
		memo[1] = 1
	}

	if _, ok := memo[n]; !ok {
		memo[n] = fib3(n-1) + fib3(n-2)
	}
	return memo[n]
}

func fib4(n int64) int64 {
	var a, b int64 = 0, 1
	for i := int64(0); i < n; i++ {
		a, b = b, a+b
	}
	return a
}
