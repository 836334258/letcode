package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root ==nil{
		return nil
	}
	var res [][]int
	queue:=[]* TreeNode{}
	queue=append(queue,root)
	for len(queue)!=0{
		r:=[]int{}
		for i:=0;i<len(queue);i++{
			node:=queue[0]
			queue=queue[1:]
			r=append(r,node.Val)
			if node.Left!=nil{
				queue=append(queue,node.Left)
			}
			if node.Right!=nil {
				queue=append(queue,node.Right)
			}
		}
		res=append(res,r)
	}

	return res
}

func main() {
	tn := &TreeNode{Val: 3, Left: &TreeNode{Left: nil, Right: nil}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15, Left: nil, Right: nil}, Right: &TreeNode{Val: 7, Left: nil, Right: nil}}}
	fmt.Println(levelOrder(tn))
}
