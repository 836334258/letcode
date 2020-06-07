package main

import (
	"fmt"
	"strconv"
)

func isHappy(n int) bool {

	temp:=[]int{}

	for  {
		tt:=strconv.Itoa(n)
		ttt:=0
		for i := 0; i < len(tt); i++ {
			t, _ := strconv.Atoi(string(tt[i]))
			ttt += t * t
		}
		n=ttt


		if n == 1 {
			return true
		}

		for t:=0;t<len(temp);t++ {
			if n==temp[t] {
				return false
			}
		}
		temp=append(temp,n)

	}
}

func main() {
	fmt.Println(isHappy(19))
}
