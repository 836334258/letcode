package main

import "fmt"

func minCount(coins []int) int {
	var ans int
	for _, v := range coins {
		if v%2 == 0 {
			ans = ans + v/2
		} else {
			ans = ans + v/2 + 1
		}
	}
	return ans
}

func main() {
	coins := []int{2, 3, 10}
	fmt.Println(minCount(coins))
}
