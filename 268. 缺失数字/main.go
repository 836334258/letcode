package main

import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {
	if len(nums)==1{
		if nums[0]==0{
			return 1
		}else{
			return 0
		}
	}

	sort.Ints(nums)
	start := nums[0]
	for i := 0; i < len(nums); i++ {
		if start == nums[i] {
			start++
			continue
		} else {
			return nums[i] -1
		}
	}

	if nums[0]>0{
		return 0
	}
	return nums[len(nums)-1]+1
}

func main() {
	a := []int{1,2}
	fmt.Println(missingNumber(a))
}
