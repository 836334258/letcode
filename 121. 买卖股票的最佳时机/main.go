package main

import "fmt"

func maxProfit(prices []int) int {
	ans:=0

	l:=len(prices)

	for i:=0;i<l-1;i++{
		for k:=i+1;k<l;k++{
			if prices[k]-prices[i]>ans  {
				ans= prices[k]-prices[i]
			}
		}
	}
	return ans
}

func main()  {
	arr:=[]int{7,6,4,3,1}
	fmt.Println(maxProfit(arr))
}
