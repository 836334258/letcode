package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	for i:=len(digits)-1;i>=0;i--{
		if digits[i]<9 {
			digits[i]++
			return digits
		}
		digits[i]=0
	}
	nums:=make([]int,len(digits)+1)
	nums[0]=1
	return nums
}

func main() {
	digits := []int{4, 3, 2, 1}
	fmt.Println(plusOne(digits))
}
