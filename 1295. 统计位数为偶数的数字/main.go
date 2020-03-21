package main

import (
	"fmt"
	"strconv"
)

func findNumbers(nums []int) int {
	i := 0
	for _, v := range nums {
		if len(strconv.Itoa(v))%2 == 0 {
			i++
		}
	}
	return i
}

func main() {
	nums := []int{12, 345, 2, 6, 7896}
	fmt.Println(findNumbers(nums))
}
