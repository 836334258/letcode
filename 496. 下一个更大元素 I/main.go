package main

import (
	"fmt"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var ans []int
	for _, v := range nums1 {
		flag := 0
		for i := 0; i < len(nums2); i++ {
			if nums2[i]==v{
				for k:=i;k<len(nums2);k++{
					if nums2[k] > v {
						ans = append(ans, nums2[k])
						flag = 1
						break
					}
				}
			}

		}

		if flag == 0 {
			ans = append(ans, -1)
		}
	}
	return ans
}

func main() {
	nums1 := []int{2,4}
	nums2 := []int{1,2,3,4}
	fmt.Println(nextGreaterElement(nums1, nums2))
}
