package services

import "fmt"

func Fib(n int) int {

	fmt.Printf("Valor N = %d\n", n)
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
