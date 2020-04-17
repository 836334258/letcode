package main

import "fmt"

func countCharacters(words []string, chars string) int {
	var charMap=make(map[uint8]int)
	var l int

	for i:=0;i<len(chars);i++{
		charMap[chars[i]]++
	}
	
	for i:=0;i<len(words);i++{
		var wordMap=make(map[uint8]int)
		flag:=true
		for j:=0;j<len(words[i]);j++{
			wordMap[words[i][j]]++
		}
		for k,_:=range wordMap{
			if wordMap[k]>charMap[k] {
				flag=false
				continue
			}
		}
		if flag {
			l+=len(words[i])
		}
	}
	return  l
}

func main()  {
	words:=[]string{"hello","world","leetcode"}
	chars:="welldonehoneyr"
	fmt.Println(countCharacters(words,chars))
}
