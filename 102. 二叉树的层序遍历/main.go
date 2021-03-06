package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) (res [][]int) {

	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil && root.Right != nil {
			res = append(res, []int{root.Left.Val, root.Right.Val})
			inOrder(root.Left)
			inOrder(root.Right)
		} else if root.Left != nil && root.Right == nil {

			res = append(res, []int{root.Left.Val})
			inOrder(root.Left)
		} else if root.Left == nil && root.Right != nil {

			res = append(res, []int{root.Right.Val})
			inOrder(root.Right)
		}
	}
	inOrder(root)
	return
}

func main() {
	node := &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}}
	fmt.Println(inorderTraversal(node))
}
