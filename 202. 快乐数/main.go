package main

import (
	"fmt"
	"strconv"
)

func isHappy(n int) bool {
	strN := strconv.Itoa(n)
	if n == 1 {
		return true
	}
	n = 0
	for i := 0; i < len(strN); i++ {
		t, _ := strconv.Atoi(string(strN[i]))
		n += t * t
	}
	return isHappy(n)
}

func main() {
	fmt.Println(isHappy(19))
}
