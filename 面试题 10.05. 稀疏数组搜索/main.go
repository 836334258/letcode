package main

import "fmt"

func findString(words []string, s string) int {
	ans:=-1
	for i:=0;i<len(words);i++{
		if words[i]==s {
			return i
		}
	}
	return ans
}

func main()  {
	words:=[]string{"at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""}
	s:="ball"
	fmt.Println(findString(words,s))
}
