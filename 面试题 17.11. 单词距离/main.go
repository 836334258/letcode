package main

import (
	"fmt"
	"math"
)

func findClosest(words []string, word1 string, word2 string) int {
	res := 1
	one, two := math.MaxInt32, math.MaxInt32
	wordsTmp := make([]int, len(words))
	for k, v := range words {
		if v == word1 {
			wordsTmp[k] = 1
		}

		if v == word2 {
			wordsTmp[k] = 2
		}
	}

	for _, v := range wordsTmp {
		if v == 1 {
			one = v
		}
		if v == 2 {
			two = v
		}
		res = int(math.Min(float64(res), float64(two-one)))
	}
	return res
}

func main() {
	words := []string{"I", "am", "a", "student", "from", "a", "university", "in", "a", "city"}
	word1 := "a"
	word2 := "student"
	fmt.Println(findClosest(words, word1, word2))
}
