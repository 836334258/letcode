package main

import "fmt"

func game(guess []int, answer []int) int {
	i := 0
	for k, _ := range guess {
		if guess[k] == answer[k] {
			i = i + 1
		}
	}
	return i
}

func main() {
	guess := []int{2,2,3}
	answer := []int{3,2,1}
	fmt.Println(game(guess, answer))
}
