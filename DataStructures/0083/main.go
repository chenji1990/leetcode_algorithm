package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	res := &ListNode{-1e9, nil}
	header := res
	for head != nil {
		if head.Val != res.Val {
			res.Next = head
			head = head.Next
			res = res.Next
			continue
		}
		head = head.Next
	}
	res.Next = nil
	return header.Next
}
