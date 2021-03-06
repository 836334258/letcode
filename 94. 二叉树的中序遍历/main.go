package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (res []int) {

	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inOrder(root.Left)
		res = append(res, root.Val)
		inOrder(root.Right)
	}
	inOrder(root)
	return
}

func main() {
	node := &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}}
	fmt.Println(inorderTraversal(node))
}
