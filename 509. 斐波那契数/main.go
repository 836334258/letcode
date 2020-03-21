package main

import "fmt"

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	return fib(N-1) + fib(N-2)
}

func main() {
	fmt.Println(fib(3))
}
