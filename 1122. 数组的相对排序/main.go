package main

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var ans []int
	sort.Ints(arr1)
	
	for _, v2 := range arr2 {
		for _, v1 := range arr1 {
			if v1 == v2 {
				ans = append(ans, v1)
			}
		}
	}

	return ans
}

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	fmt.Println(relativeSortArray(arr1, arr2))
}
