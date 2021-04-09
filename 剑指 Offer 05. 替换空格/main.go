package main

import (
	"fmt"
	"strings"
)

func replaceSpace(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

func main() {
	s := "We are happy."
	fmt.Println(replaceSpace(s))
}
