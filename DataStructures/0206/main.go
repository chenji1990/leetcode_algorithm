package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	res := &ListNode{0, nil}
	var temp *ListNode = nil
	for head != nil {
		temp = head.Next
		head.Next = res.Next
		res.Next = head
		head = temp
	}
	return res.Next
}
