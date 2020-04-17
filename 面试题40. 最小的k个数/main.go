package main

import (
	"fmt"
	"sort"
)

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

func main()  {
	arr:=[]int{0,1,2,1}
	k:=1
	fmt.Println(getLeastNumbers(arr,k))
}
