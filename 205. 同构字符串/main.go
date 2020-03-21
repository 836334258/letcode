package main

import (
	"fmt"
	"strings"
)

func isIsomorphic(s string, t string) bool {
	for k, v := range s {
		if strings.Index(s, string(v)) != strings.Index(t, string(t[k])) {
			return false
		}
	}
	return true
}

func main() {
	s := "foo"
	t := "bar"
	fmt.Println(isIsomorphic(s, t))
}
