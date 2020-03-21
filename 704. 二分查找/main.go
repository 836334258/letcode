package main

import "fmt"

func search(nums []int, target int) int {
	ans := -1
	for k, v := range nums {
		if v == target {
			ans = k
		}
	}
	return ans
}

func main() {
	nums := []int{-1,0,3,5,9,12}
	target := 2
	fmt.Println(search(nums, target))
}
