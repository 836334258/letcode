package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	var tmp []int
	for i := 0; i < len(nums1); i++ {
		for k := 0; k < len(nums2); k++ {
			if nums1[i] == nums2[k] {
				sign := 0
				for _, v := range tmp {
					if v == nums1[i] {
						sign = 1
					}
				}
				if sign == 0 {
					tmp = append(tmp, nums1[i])
				}
			}
		}
	}
	return tmp
}

func main() {
	nums1 := []int{4,9,5}
	nums2 := []int{9,4,9,8,4}
	fmt.Println(intersection(nums1, nums2))
}
