package main

import "fmt"

func sortArrayByParity(A []int) []int {
	even := []int{}
	odd := []int{}

	for _, v := range A {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

	return append(even, odd...)
}

func main() {
	A := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(A))
}
