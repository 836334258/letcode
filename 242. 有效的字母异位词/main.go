package main

import (
	"fmt"
	"sort"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s1 := []byte(s)
	t1 := []byte(t)

	sort.Slice(s1, func(i, j int) bool {
		return s1[i] <s1[j]
	})

	sort.Slice(t1, func(i, j int) bool {
		return t1[i] < t1[j]
	})

	for k,v:=range s1{
		if v!=t1[k]{
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
