package main

import "fmt"

func selfDividingNumbers(left int, right int) []int {
	var ans []int
	for i := left; i <= right; i++ {
		tmp := i
		for tmp > 0 {
			tmp2 := tmp % 10
			if tmp2==0 {
				break
			}
			if tmp2!=0&&i%tmp2 != 0 {
				break
			}
			tmp /= 10
		}
		if tmp == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	left := 1
	right := 22
	fmt.Println(selfDividingNumbers(left, right))
}
