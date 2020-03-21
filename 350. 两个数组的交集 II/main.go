package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var ans []int
	i := 0
	k := 0
	for i <len(nums1) && k < len(nums2) {
		if nums1[i] > nums2[k] {
			k++
		} else if nums1[i] < nums2[k] {
			i++
		} else {
			ans = append(ans, nums1[i])
			i++
			k++
		}
	}
	return ans
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersect(nums1, nums2))
}
