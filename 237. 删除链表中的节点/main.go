package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val=node.Next.Val
	node.Next=node.Next.Next
}

func main()  {
	node:=&ListNode{Val:1,Next:&ListNode{Val:0,Next:&ListNode{Val:1,Next:nil}}}
	for node!=nil{
		fmt.Println(node.Val)
		node=node.Next
	}
}
