package main

import (
	"fmt"
	"strings"
	"unicode"
)

func letterCasePermutation(S string) []string {
	s := []string{}
	for k, v := range S {
		//判断是不是字母
		if unicode.IsLetter(v) {
			s = append(s, S[:k]+strings.ToLower(string(v))+S[k+1:])
			s = append(s, S[:k]+strings.ToUpper(string(v))+S[k+1:])
		}
	}
	if len(s) == 0 {
		return append(s, S)
	}
	return s
}

func main() {
	S := "a1b2"
	fmt.Println(letterCasePermutation(S))
}
