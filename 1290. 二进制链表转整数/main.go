package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	num:=0
	for head!=nil{
		num=num<<1+head.Val
		head=head.Next
	}
	return num

}


func main()  {
	ln:=&ListNode{Val:1,Next:&ListNode{Val:0,Next:&ListNode{Val:1,Next:nil}}}


	fmt.Println(getDecimalValue(ln))

}
