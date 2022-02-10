package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}

	node1 := head
	node2 := head.Next
	for node2 != nil {
		if node2.Val == val {
			node1.Next = node2.Next
			break
		}
		node1 = node1.Next
		node2 = node2.Next
	}
	return head
}
