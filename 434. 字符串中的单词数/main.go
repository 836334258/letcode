package main

import (
	"fmt"
	"strings"
)

func countSegments(s string) int {
	n := 0
	if strings.Count(s, " ") == len(s) {
		return 0
	}
	for k, _ := range s {
		if k == 0 && string(s[k]) == " " && string(s[k+1]) != " " {
			n++
		} else if k == 0 && string(s[k]) == " " && string(s[k+1]) == " " {
			continue
		} else if k == 0 || string(s[k-1]) == " " && string(s[k]) != " " {
			n++
		}
	}
	return n
}

func main() {
	str := "Hello, my name is John"
	fmt.Println(countSegments(str))
}
