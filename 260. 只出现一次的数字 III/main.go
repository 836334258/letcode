package main

import "fmt"

func singleNumber(nums []int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	var tmp []int
	for k, v := range m {
		if v == 1 {
			tmp = append(tmp, k)
		}
	}

	return tmp
}

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	fmt.Println(singleNumber(nums))
}
