package main

import "fmt"

func sumOfUnique(nums []int) int {
	mp := map[int]int{}
	for _, v := range nums {
		mp[v] = mp[v] + 1
	}
	ans := 0
	for k, v := range mp {
		if v == 1 {
			ans = ans + k
		}
	}

	return ans
}

func main() {
	nums := []int{1, 1, 1, 1, 1}
	fmt.Println(sumOfUnique(nums))
}
