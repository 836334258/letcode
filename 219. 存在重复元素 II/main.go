package main

import (
	"fmt"
	"math"
	"strconv"
)

//超时
func containsNearbyDuplicate1(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i+1; j <= len(nums)-1; j++ {
			if nums[i] == nums[j] && int(math.Abs(float64(i-j))) <= k {
				return true
			}
		}
	}
	return false
}

func containsNearbyDuplicate(nums []int, k int) bool {
	strMap:=make(map[string]int)
	for i:=0;i<len(nums);i++{
		strMap[strconv.Itoa(nums[i])]++
		if strMap[strconv.Itoa(nums[i])]>=2 {
			for j:=0;j<i;j++{
				if nums[i]==nums[j] {
					if i-j<=k {
						return true
					}
				}
			}
		}
	}
	return false
}

func main() {
	nums := []int{1,0,1,1}
	k := 1
	fmt.Println(containsNearbyDuplicate(nums, k))
}
