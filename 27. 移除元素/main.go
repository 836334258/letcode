package main

import "fmt"

func removeElement(nums []int, val int) int {
	index:=0
	for ;index<len(nums);{
		if nums[index]==val {
			nums=append(nums[:index],nums[index+1:]...)
			continue
		}
		index++
	}
	return len(nums)
}

func main()  {
	nums:=[]int{3,2,2,3}
	val:=3
	fmt.Println(removeElement(nums,val))
}
