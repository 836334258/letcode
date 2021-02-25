package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	var ansArr [][]int
	for k1, _ := range nums {
		for k2, _ := range nums {
			if nums[k1] == nums[k2] && k1 < k2 {
				ansArr = append(ansArr, []int{k1, k2})
			}
		}
	}
	return len(ansArr)
}

func main() {
	nums := []int{1, 1, 1, 1}
	fmt.Println(numIdenticalPairs(nums))
}
