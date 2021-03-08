package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	var res int
	var fn func(root *TreeNode)
	if root == nil {
		return 0
	}
	fn = func(root *TreeNode) {
		if root != nil {
			res++
		}
		if root.Left != nil {
			fn(root.Left)
		}
		if root.Right != nil {
			fn(root.Right)
		}
	}

	fn(root)

	return res
}

func main() {
	fmt.Println(countNodes(&TreeNode{Val: 1, Left: nil, Right: nil}))
}
