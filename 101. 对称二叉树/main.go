package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root,root)
}

func isMirror(t1 *TreeNode,t2 * TreeNode) bool {
	if t1==nil && t2==nil {
		return true
	}
	if t1==nil||t2==nil {
		return false
	}

	return t1.Val==t2.Val && isMirror(t1.Left,t2.Right) && isMirror(t1.Right,t2.Left)
}

func main() {
	t:=&TreeNode{Val:1,Left:&TreeNode{Val:2,Left:&TreeNode{Val:3},Right:&TreeNode{Val:4}},Right:
	&TreeNode{Val:2,Left:&TreeNode{Val:4},Right:&TreeNode{Val:3}}}
	fmt.Println(isSymmetric(t))
}
