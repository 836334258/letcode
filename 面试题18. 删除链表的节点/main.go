package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}

	for cur := dummy; cur != nil && cur.Next != nil; cur = cur.Next {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		}
	}
	return dummy.Next
}

func main() {
	head := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 9, Next: nil}}}}
	//for head.Next != nil {
	//	fmt.Println(head.Val)
	//	head.Next = head.Next.Next
	//}
	fmt.Println(deleteNode(head, 5))
}
