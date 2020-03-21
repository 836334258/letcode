package main

import "fmt"

func checkRecord(s string) bool {
	a:=0
	b:=0
	for _,v:=range s{
		if string(v)=="A"{
			a++
			if a>1 {
				return false
			}
		}
		if string(v)=="L" {
			b++
			if b>2{
				return false
			}
		}else {
			b=0
		}
	}
	return  true
}

func main()  {
	s:="LALL"
	fmt.Println(checkRecord(s))
}
