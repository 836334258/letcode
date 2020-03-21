package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3,8,-10,23,19,-4,-14,27}
	sort.Ints(arr)
	left := 0
	right := 1
	c := arr[right] - left
	var res [100][2]int
	for right < len(arr) {
		k := arr[right] - arr[left]
		if k < c {
			res[left][0] = arr[left]
			res[left][1] = arr[right]
		}
		left++
		right++
	}

	fmt.Println(res)
}
