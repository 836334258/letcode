package main

import "fmt"

func generate(numRows int) [][]int {
	var ans [][]int
	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			k++
			ans[i][j] = k
			fmt.Printf("a[%d][%d]=%d", i, j, ans[i][j])
		}
		fmt.Println("\n")
	}

	return ans

}

func main() {
	fmt.Println(generate(5))
}
