package main

import "fmt"

func numberOfSteps(num int) int {
	step:=0
	for num>0{
		if num%2==0{
			num/=2
		}else{
			num-=1
		}
		step++
	}
	return  step
}

func main() {
	num := 123
	fmt.Println(numberOfSteps(num))
}
