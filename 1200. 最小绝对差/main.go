package main

import (
	"fmt"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	fmt.Println(arr)
	var aa [][]int

	min := arr[1] - arr[0]
	for i := 0; i < len(arr)-1; i++ {
		if min == arr[i+1]-arr[i] {
			aa = append(aa, []int{arr[i], arr[i+1]})
		} else if min > arr[i+1]-arr[i] {
			min = arr[i+1] - arr[i]
			aa = [][]int{}
			aa = append(aa, []int{arr[i], arr[i+1]})
		}
	}

	return aa
}

func main() {
	aa := []int{40, 11, 26, 27, -20}
	fmt.Println(minimumAbsDifference(aa))
}
