package main

import (
	"fmt"
	"strings"
)

func uncommonFromSentences(A string, B string) []string {
	arrA := strings.Split(A, " ")
	arrB := strings.Split(B, " ")
	arrA = append(arrA, arrB...)
	var answer []string
	for _, v1 := range arrA {
		i := 0
		for _, v2 := range arrA {
			if v1 == v2 {
				i++
			}
		}
		if i == 1 {
			answer = append(answer, v1)
		}
	}
	return answer
}

func main() {
	A := "this apple is sweet"
	B := "this apple is sour"
	fmt.Println(uncommonFromSentences(A, B))
}
