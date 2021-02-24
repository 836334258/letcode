package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, v1 := range accounts {
		tmp := 0
		for _, v2 := range v1 {
			tmp += v2
		}
		if max < tmp {
			max = tmp
		}
	}

	return max
}

func main() {
	accounts := [][]int{{1, 5}, {7, 3}, {3, 5}}
	fmt.Println(maximumWealth(accounts))
}
