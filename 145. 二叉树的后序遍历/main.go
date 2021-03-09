package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) (res []int) {
	var f func(root *TreeNode)

	f = func(root *TreeNode) {
		if root == nil {
			return
		}

		f(root.Left)
		f(root.Right)
		res = append(res, root.Val)
	}

	f(root)
	return
}

func main() {

}
