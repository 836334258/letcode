package main

import (
	"fmt"
	"strings"
)

func replaceWords(dict []string, sentence string) string {
	sentenceArr:=strings.Split(sentence," ")
	for k,v:=range sentenceArr{ //"aadsfasf absbs bbab cadsfafs"
		for d1,v1:=range dict{  //["a","b","c"]
			if strings.Index(v,v1)==0 {
				sentenceArr[k]=dict[d1]
				break
			}
		}
	}
	return  strings.Join(sentenceArr," ")
}

func main()  {
	dict:=[]string{"a", "b", "c"}
	scnetence:="aadsfasf absbs bbab cadsfafs"
	fmt.Println(replaceWords(dict,scnetence))
}
