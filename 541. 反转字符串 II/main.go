package main

import (
	"fmt"
)

func reverseStr(s string, k int) string {
	var tmp string
	sLeft := s[:k]
	fmt.Println(sLeft)
	for i := k-1; i >= 0; i-- {
		tmp += string(sLeft[i])
	}
	sRight := s[k:]
	return tmp + sRight
}
func main() {
	s := "abcdefg"
	k := 2
	fmt.Println(reverseStr(s, k))
}
