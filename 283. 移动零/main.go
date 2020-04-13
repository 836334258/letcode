package main

import "fmt"

func moveZeroes(nums []int)  {

	next:=0
	l:=len(nums)
	for i:=0;i<l;i++{
		if nums[i]!=0 {
			nums[next]=nums[i]
			next++
		}
	}
	for next<l{
		nums[next]=0
		next++
	}
	return
}

func main()  {
	nums:=[]int{0,1,0,3,12}

	moveZeroes(nums)

	fmt.Println(nums)
}
