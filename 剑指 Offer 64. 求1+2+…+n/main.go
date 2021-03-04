package main

import "fmt"

func sumNums(n int) int {
	if n == 0 {
		return 0
	}

	return n + sumNums(n-1)
}

func main() {
	n := 9
	fmt.Println(sumNums(n))
}
