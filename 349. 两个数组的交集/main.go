package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	return  nums1 & nums2
}

func main()  {
	nums1:=[]int{1,2,2,1}
	nums2:=[]int{2,2}
	fmt.Println(intersection(nums1,nums2))
}
