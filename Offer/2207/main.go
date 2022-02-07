package main

import "container/list"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	stack := list.New()
	for head != nil {
		stack.PushBack(head.Val)
		head = head.Next
	}
	res := []int{}
	for stack.Len() > 0 {
		res = append(res, stack.Back().Value.(int))
		stack.Remove(stack.Back())
	}
	return res
}
