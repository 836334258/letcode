package main

import (
	"fmt"
	"sort"
)

func frequencySort(nums []int) []int {
	var tmp map[int]int
	tmp = make(map[int]int)
	for _, v := range nums {
		tmp[v]++
	}

	fmt.Println(tmp)

	sort.Slice()
	return nums

}

func main() {
	nums := []int{1, 1, 2, 2, 2, 3}

	frequencySort(nums)
}
