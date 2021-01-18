package main

import "fmt"

func hammingWeight(num uint32) int {
	var t int
	var tmp = num
	for tmp != 0 {
		if tmp&1 == 1 {
			t++
		}
		tmp = tmp >> 1
	}
	return t
}

func main() {
	var num uint32 = 11111111111111111111111111111101
	fmt.Println(hammingWeight(num))
}
