package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	res := &ListNode{0, nil}
	header := res

	for head != nil {
		if head.Val != val {
			res.Next = head
			res = res.Next
			head = head.Next
			continue
		}
		head = head.Next
	}
	res.Next = nil
	return header.Next
}
