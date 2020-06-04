package main

import (
	"fmt"
)

func reverseString(s []byte)  {
	for i:=0;i<(len(s))/2;i++ {
		tmp:=s[i]
		s[i]=s[len(s)-1-i]
		s[len(s)-i-1]=tmp

	}
}

func main()  {
	arr:=[]byte{'h','e','l','l','o'}
	reverseString(arr)

	fmt.Println((string(arr)))
}
