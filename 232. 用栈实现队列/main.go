package main
//
//type MyQueue struct {
//	stackPush []int
//	stackPop []int
//}
//
//
///** Initialize your data structure here. */
//func Constructor() MyQueue {
//	return MyQueue{}
//}
//
//
///** Push element x to the back of queue. */
//func (this *MyQueue) Push(x int)  {
//	this.stackPush=append(this.stackPush,x)
//	if len(this.stackPop)<=0{
//		for _,val:=range this.stackPush{
//			this.stackPop=append(this.stackPop,val)
//		}
//		this.stackPush=nil
//	}
//}
//
//
///** Removes the element from in front of queue and returns that element. */
//func (this *MyQueue) Pop() int {
//	if (len(this.stackPop)<=0&&len(this.stackPush)<=0)){
//		return 0
//	}
//
//}
//
//
///** Get the front element. */
//func (this *MyQueue) Peek() int {
//
//}
//
//
///** Returns whether the queue is empty. */
//func (this *MyQueue) Empty() bool {
//
//}
//
//
///**
// * Your MyQueue object will be instantiated and called as such:
// * obj := Constructor();
// * obj.Push(x);
// * param_2 := obj.Pop();
// * param_3 := obj.Peek();
// * param_4 := obj.Empty();
// */
//
//func main()  {
//
//}
