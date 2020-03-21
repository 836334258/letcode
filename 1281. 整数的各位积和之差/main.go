package main

import (
	"fmt"
	"strconv"
)
func subtractProductAndSum(n int) int {
	strN:=strconv.Itoa(n)
	i:=len(strN)
	ji:=1
	he:=0
	for i>0{
		aa:=n%10
		n/=10
		ji*=aa
		he+=aa
		i--
	}
	return ji-he
}


func main()  {
	fmt.Println(subtractProductAndSum(234))
}
