package main

import (
	"fmt"
)

func findSpecialInteger(arr []int) int {
	var m = make(map[int]int)
	for k, _ := range arr {
		m[arr[k]]++
	}
	max := 0
	index := 0
	for k, v := range m {
		if v > max {
			max, index = v, k
		}
	}
	return index
}

func main() {
	arr := []int{1}
	fmt.Println(findSpecialInteger(arr))
}
