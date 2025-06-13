package main

import "fmt"

type MinStack struct {
	dataStack []int
	minStack  []int
}

func Constructor() MinStack {
	return MinStack{
		dataStack: []int{},
		minStack:  []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.dataStack = append(this.dataStack, val)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
	} else {
		if this.minStack[len(this.minStack)-1] > val {
			this.minStack = append(this.minStack, val)
		} else {
			this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.dataStack) > 0 {
		this.dataStack = this.dataStack[:len(this.dataStack)-1]
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.dataStack) > 0 {
		return this.dataStack[len(this.dataStack)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.minStack) > 0 {
		return this.minStack[len(this.minStack)-1]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	val := 1
	obj.Push(val)
	obj.Pop()
	param3 := obj.Top()
	param4 := obj.GetMin()
	fmt.Printf("param3:%d, param4:%d\n", param3, param4)
}
