package main

import "fmt"

func calculate(s string) int {
	x := 1
	y := 0

	for _, v := range s {
		if string(v) == "A" {
			x = 2*x + y
		} else {
			y = 2*y + x
		}
	}
	return x + y
}

func main() {
	s := "AB"
	fmt.Println(calculate(s))
}
