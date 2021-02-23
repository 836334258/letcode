package main

import "fmt"

func xorOperation(n int, start int) int {
	ans := 0
	i := 0
	for n > 0 {
		ans = ans ^ (start + 2*i)
		i = i + 1
		n = n - 1
	}
	return ans
}

func main() {
	fmt.Println(xorOperation(5, 0))
}
