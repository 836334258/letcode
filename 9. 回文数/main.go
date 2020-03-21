package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	sx := strconv.Itoa(x)
	l := len(sx)-1
	for i := 0; i <l; i++ {
		if sx[i] != sx[l-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(121))
}
