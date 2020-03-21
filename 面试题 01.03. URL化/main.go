package main

import (
	"fmt"
	"strings"
)

func replaceSpaces(S string, length int) string {
	//var ans string
	//sLen := len(S)
	//if sLen < length {
	//	length = sLen
	//}
	//fmt.Println(length)
	//for i := 0; i < length; i++ {
	//	if S[i] == ' ' {
	//		ans += "%20"
	//		continue
	//	}
	//	ans += string(S[i])
	//}
	//return ans
	return strings.Replace(S[:length]," ","%20",-1)
}

func main() {
	s := "               "
	len := 5
	fmt.Println(replaceSpaces(s, len))
}
