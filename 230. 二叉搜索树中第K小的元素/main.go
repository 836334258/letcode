package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	arr := []int{}
	find(root, &arr)
	return arr[k-1]

}

func find(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	find(root.Left, arr)
	*arr = append(*arr, root.Val)
	find(root.Right, arr)
	return
}

func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: nil, Right: nil}}, Right: &TreeNode{Val: 4, Left: nil, Right: nil}}
	fmt.Println(kthSmallest(root, 1))
}
