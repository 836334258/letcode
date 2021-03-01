package main

import "fmt"

func decompressRLElist(nums []int) []int {
	var ans []int
	for i := 0; i < len(nums); i += 2 {
		for k := 0; k < nums[i]; k++ {
			ans = append(ans, nums[i+1])
		}
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(decompressRLElist(nums))
}
