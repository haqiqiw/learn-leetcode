package main

import "fmt"

type MinStack struct {
	items []int
	mins  []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.items) == 0 {
		this.mins = append(this.mins, val)
	} else {
		this.mins = append(this.mins, min(val, this.mins[len(this.mins)-1]))
	}

	this.items = append(this.items, val)
}

func (this *MinStack) Pop() {
	this.mins = this.mins[:len(this.mins)-1]
	this.items = this.items[:len(this.items)-1]
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}

func (this *MinStack) Print() {
	fmt.Println(this.items)
	fmt.Println(this.mins)
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
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Print()
	minStack.Push(0)
	minStack.Print()
	minStack.Push(-3)
	minStack.Print()
	fmt.Println(minStack.GetMin()) // return -3
	minStack.Pop()
	minStack.Print()
	fmt.Println(minStack.Top())    // return 0
	fmt.Println(minStack.GetMin()) // return -2
}
