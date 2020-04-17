package main

import "fmt"

func majorityElement(nums []int) int {
	var m = make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	l := len(nums)
	for k, v := range m {

		if v > l/2 {
			return k
		}
	}
	return -1
}

func main() {
	nums := []int{3,2,3}
	fmt.Println(majorityElement(nums))
}
