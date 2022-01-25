package main

import (
	"container/list"
)

type MinStack struct {
	stack    *list.List
	minStack *list.List
}

func Constructor() MinStack {
	return MinStack{
		stack:    list.New(),
		minStack: list.New(),
	}
}

func (this *MinStack) Push(x int) {
	if this.stack.Len() == 0 {
		this.minStack.PushBack(x)
	} else {
		this.minStack.PushBack(min(x, this.minStack.Back().Value.(int)))
	}
	this.stack.PushBack(x)
}

func (this *MinStack) Pop() {
	if this.stack.Len() > 0 {
		this.stack.Remove(this.stack.Back())
		this.minStack.Remove(this.minStack.Back())
	}
}

func (this *MinStack) Top() int {
	return this.stack.Back().Value.(int)
}

func (this *MinStack) Min() int {
	return this.minStack.Back().Value.(int)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
