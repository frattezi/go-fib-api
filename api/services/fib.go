package services

import "fmt"

func Fib(n int) int {

	fmt.Printf("Valor N = %d\n", n)
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// package services

// Memory fib implementation
// func FibRedis(n int, lookup map[int]int) int {
// 	if n == 1 || n == 0 {
// 		lookup[n] = n
// 	}

// 	if _, ok := lookup[n]; !ok {
// 		lookup[n] = FibRedis(n-1, lookup) + FibRedis(n-2, lookup)
// 	}

// 	return lookup[n]
// }
