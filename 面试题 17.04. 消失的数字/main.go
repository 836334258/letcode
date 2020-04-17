package main

import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {
	sort.Ints(nums)
	for k,v:=range nums{
		if k!=v {
			return k
		}
		if v==len(nums)-1 {
			return len(nums)
		}
	}
	return 1
}

func main()  {
	nums:=[]int{0,1}
	fmt.Println(missingNumber(nums))
}
