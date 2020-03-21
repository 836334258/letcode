package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	var tmp []int
	for i:=0;i<len(arr);i++{
		tmp[arr[i]]++
	}
	fmt.Println(tmp)
	return  true
}

func main() {
	arr := []int{1, 2, 2, 1, 1, 3}
	fmt.Println(uniqueOccurrences(arr))
}
