package main

import "fmt"

func isValid(s string) bool {
	l := len(s)
	if l == 1 {
		return false
	}
	for i := 0; i < l; i += 2 {
		switch string(s[i]) {
		case "(":
			if string(s[i+1]) != ")" && string(s[l-i-1]) != ")" {
				return false
			}
		case "{":
			if string(s[i+1]) != "}" && string(s[l-i-1]) != "}" {
				return false
			}
		case "[":
			if string(s[i+1]) != "]" && string(s[l-i]-1) != "]" {
				return false
			}
		}
	}

	return true
}

func main() {
	s := "{[]}"
	fmt.Println(isValid(s))
}
