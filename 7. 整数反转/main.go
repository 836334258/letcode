package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	if x >= math.MaxInt32 || x < (-math.MaxInt32) {
		return 0
	}
	var ans string
	if x >= 0 {
		for x != 0 {
			tmp := x % 10
			if tmp != '0' {
				ans += strconv.Itoa(tmp)
			}
			x /= 10
		}

		k, _ := strconv.Atoi(ans)
		if k >= math.MaxInt32 || k < (-math.MaxInt32) {
			return 0
		}
		return k
	} else {
		xStr := strconv.Itoa(x)[1:]
		x, _ := strconv.Atoi(xStr)
		for x != 0 {
			tmp := x % 10
			if tmp != '0' {
				ans += strconv.Itoa(tmp)
			}
			x /= 10
		}
		ans = "-" + ans
		k, _ := strconv.Atoi(ans)
		if k >= math.MaxInt32 || k < (-math.MaxInt32) {
			return 0
		}
		return k
	}

}

func main() {
	fmt.Println(reverse(1534236469))
}
