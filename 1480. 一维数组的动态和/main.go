package main

import "fmt"

func runningSum(nums []int) []int {
	var tmp int
	var tmpArr []int

	for _, v := range nums {
		tmp += v
		tmpArr = append(tmpArr, tmp)
	}

	return tmpArr
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(runningSum(nums))
}
