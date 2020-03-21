package main

import "fmt"

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) == 1 || len(coordinates) == 2 {
		return true
	}
	//斜率
	k := (coordinates[1][1] - coordinates[0][1]) / (coordinates[1][0] - coordinates[0][0])
	for i := 0; i < len(coordinates)-1; i++ {
		if(coordinates[i+1][0]-coordinates[i][0])==0|| (coordinates[i+1][1]/coordinates[i][1])/(coordinates[i+1][0]-coordinates[i][0]) != k  {
			return false
		}
	}
	return true
}

func main() {
	a := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	fmt.Println(checkStraightLine(a))
}
