package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	var s string
	for k, v := range word1 {
		if k == len(word2) {
			break
		}
		s = s + string(v)
		s = s + string(word2[k])
	}

	if len(word2) > len(word1) {
		s = s + word2[len(word1):]
	}

	if len(word2) < len(word1) {
		s = s + word1[len(word2):]
	}
	return s
}

func main() {
	word1 := "abcd"
	word2 := "pq"
	fmt.Println(mergeAlternately(word1, word2))
}
