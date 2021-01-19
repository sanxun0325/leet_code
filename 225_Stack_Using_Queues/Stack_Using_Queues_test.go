package _25_Stack_Using_Queues

import "container/list"

type MyStack struct {
	list *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{list: list.New()}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.list.PushFront(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	v := this.list.Front()
	this.list.Remove(v)
	return v.Value.(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.list.Front().Value.(int)
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.list.Front() == nil {
		return true
	}
	return false
}
