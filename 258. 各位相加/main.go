package main

import (
	"fmt"
	"strconv"
)

func addDigits(num int) int {
	if len(strconv.Itoa(num)) == 1 {
		return num
	}
	k := 0
	for num > 0 {
		k += num % 10
		num /= 10
	}
	return addDigits(k)
}

func main() {
	fmt.Println(addDigits(38))
}
