package main

import "fmt"

func shuffle(nums []int, n int) []int {
	var a []int
	var b []int

	for i := 0; i < len(nums); i++ {
		if i < n {
			a = append(a, nums[i])
		} else {
			b = append(b, nums[i])
		}
	}

	var c []int

	for i := 0; i < len(a); i++ {
		c = append(c, a[i])
		c = append(c, b[i])
	}

	return c
}

func main() {
	nums := []int{2, 5, 1, 3, 4, 7}
	n := 3

	fmt.Println(shuffle(nums, n))
}
