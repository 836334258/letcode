package main

import "fmt"

func maxSubArray(nums []int) int {
	b:=nums[0]
	sum:=b
	for i:=1;i<len(nums);i++{
		if b<0 {
			b=nums[i]
		}else {
			b+=nums[i]
		}
		if b>sum {
			sum=b
		}
	}
	return sum
}
func main()  {
	nums:=[]int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(nums))
}

