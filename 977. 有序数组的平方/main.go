package main

import (
	"fmt"
	"sort"
)

func sortedSquares(nums []int) []int {
	var tmp []int

	for _, v := range nums {
		tmp = append(tmp, v*v)
	}

	sort.Ints(tmp)
	return tmp
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}
