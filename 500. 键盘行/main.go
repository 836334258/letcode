package main

import (
	"fmt"
	"strings"
)

func findWords(words []string) []string {
	firstRow := "qwertyuiopQWERTYUIOP"
	secondRow := "asdfghjklASDFGHJKL"
	thirdRow := "zxcvbnmZXCVBNMZ"
	var answer []string

	for _, v := range words {
		line := firstRow
		flag := true
		if strings.Contains(secondRow, string(v[0])) {
			line = secondRow
		} else if strings.Contains(thirdRow, string(v[0])) {
			line = thirdRow
		}
		for _, c := range v {
			if !strings.Contains(line, string(c)) {
				flag = false
				break
			}
		}
		if flag {
			answer = append(answer, v)
		}
	}
	return answer
}

func main() {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}


	fmt.Println(findWords(words))

}
