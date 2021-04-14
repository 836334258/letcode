package main

import (
	"fmt"
)

func restoreString(s string, indices []int) string {
	ss := make([]rune, len(indices))
	for k, v := range s {
		ss[indices[k]] = v
	}
	return string(ss)
}

func main() {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	fmt.Println(restoreString(s, indices))
}
