package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var root *ListNode = &ListNode{0, nil}
	var temp *ListNode

	for head != nil {
		temp = head.Next
		head.Next = root.Next
		root.Next = head
		head = temp
	}
	return root.Next
}
