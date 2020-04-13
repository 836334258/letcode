package main

import (
	"fmt"
	"strconv"
	"strings"
)

func plusOne(digits []int) []int {
	var tmp []string

	for _, k := range digits {
		tmp = append(tmp, strconv.Itoa(k))
	}
	tmp1, _ := strconv.Atoi(strings.Join(tmp, ""))
	tmp1++
	tmp2 := strconv.Itoa(tmp1)

	var ans []int
	for _, v := range tmp2 {
		kk, _ := strconv.Atoi(string(v))
		ans = append(ans, kk)
	}
	return ans
}

func main() {
	digits := []int{4, 3, 2, 1}
	fmt.Println(plusOne(digits))
}
