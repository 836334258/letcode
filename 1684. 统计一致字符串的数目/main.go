package main

import (
	"fmt"
	"strings"
)

func countConsistentStrings(allowed string, words []string) int {
	ans := 0
	for _, v := range words {
		for k, char := range v {
			if !strings.Contains(allowed, string(char)) {
				break
			}

			if k == len(v)-1 {
				ans = ans + 1
			}
		}
	}
	return ans
}

func main() {
	allowed := "abc"
	words := []string{"a", "b", "c", "ab", "ac", "bc", "abc"}
	fmt.Println(countConsistentStrings(allowed, words))
}
