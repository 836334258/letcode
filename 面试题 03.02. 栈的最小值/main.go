package main

type MinStack struct {
	arr []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.arr = append(this.arr, x)
}

func (this *MinStack) Pop() {
	this.arr = this.arr[:len(this.arr)-1]
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.arr) == 0 {
		return 0
	}

	arrMin := this.arr[0]
	for _, v := range this.arr {
		if v < arrMin {
			arrMin = v
		}
	}
	return arrMin
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)

	print(obj.GetMin())
	obj.Pop()

	print(obj.Top())
	print(obj.GetMin())
}
