package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n==0{
		return false
	}
	for n%4 == 0 {
		n /= 4
	}
	return n == 1
}

func main() {
	fmt.Println(isPowerOfTwo(5))
}
