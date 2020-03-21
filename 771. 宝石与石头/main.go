package main

import (
	"fmt"
	"strings"
)

func numJewelsInStones(J string, S string) int {
	ans:=0
	for _,v:=range J{
		ans+=strings.Count(S,string(v))
	}
	return ans
}

func main()  {
	fmt.Println(numJewelsInStones("aA","aAAbbbb"))

}
