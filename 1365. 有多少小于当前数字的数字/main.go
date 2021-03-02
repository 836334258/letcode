package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	ans := []int{}

	for _, v1 := range nums {
		var tmp int
		for _, v2 := range nums {
			if v1 > v2 {
				tmp++
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}

func main() {
	nums := []int{8, 1, 2, 2, 3}
	fmt.Println(smallerNumbersThanCurrent(nums))
}
