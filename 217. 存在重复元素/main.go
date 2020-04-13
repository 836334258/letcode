package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	if len(nums)==1 {
		return false
	}
	for i:=0;i<len(nums)-1;i++{
		if nums[i]==nums[i+1] {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{1,1,1,3,3,4,3,2,4,2}
	fmt.Println(containsDuplicate(arr))
}
