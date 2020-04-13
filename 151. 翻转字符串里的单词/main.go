package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	trimStr:=strings.TrimSpace(s)
	trimArr:=strings.Split(trimStr," ")
	var ansArr []string
	for i:=len(trimArr)-1;i>=0;i--{
		if strings.Trim(trimArr[i]," ")=="" {
			continue
		}
		ansArr=append(ansArr,trimArr[i])
	}
	return strings.Join(ansArr," ")
}

func main()  {
	s:="a good   example"
	fmt.Println(reverseWords(s))
}
