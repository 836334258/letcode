package main

import "fmt"

func singleNumber(nums []int) int {
	var i = nums[0]
	for k := 1; k < len(nums); k++ {
		i = i ^ nums[k]
	}
	return i
}

func main() {
	arr := []int{2, 2, 1}
	fmt.Println(singleNumber(arr))
}
