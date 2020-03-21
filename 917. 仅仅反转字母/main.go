package main

import (
	"fmt"
	"unicode"
)

func reverseOnlyLetters(S string) string {
	var s string
	for _, k := range S {
		if unicode.IsLetter(k) {
			s = string(k) + s
		}
	}
	for k, v := range S {
		if !unicode.IsLetter(v) {
			s = s[:k] + string(v) + s[k:]
		}
	}

	return s
}

func main() {
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj"))
}
