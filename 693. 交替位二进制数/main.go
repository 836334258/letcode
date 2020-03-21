package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hasAlternatingBits(n int) bool {

	str:=""
	for n>0{
		str+=strconv.Itoa(n%2)
		n/=2
	}
	return ! (strings.Contains(str,"00") || strings.Contains(str,"11"))
}

func main()  {
	n:=10
	fmt.Println(hasAlternatingBits(n))
}
