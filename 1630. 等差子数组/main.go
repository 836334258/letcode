package main

import (
	"fmt"
	"sort"
)

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	var ans []bool

	for i := 0; i < len(l); i++ {
		var tmpArr []int

		for g := l[i]; g <= r[i]; g++ {
			tmpArr = append(tmpArr, nums[g])
		}
		tempBool := true
		sort.Ints(tmpArr)
		fmt.Println(tmpArr)
		for k := 0; k < len(tmpArr)-1; k++ {
			tmp := tmpArr[1] - tmpArr[0]
			if tmpArr[k+1]-tmpArr[k] != tmp {
				//fmt.Println(tmpArr)
				//fmt.Println(k)
				tempBool = false
			}
		}

		ans = append(ans, tempBool)
	}

	return ans
}

func main() {
	nums := []int{4, 6, 5, 9, 3, 7}
	l := []int{0, 0, 2}
	r := []int{2, 3, 5}

	fmt.Println(checkArithmeticSubarrays(nums, l, r))
}
