package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	length := checkSN(head)
	for i := 0; i < length-k; i++ {
		head = head.Next
	}
	return head
}

func checkSN(node *ListNode) int {
	if node == nil {
		return 0
	}
	if node.Next == nil {
		return 1
	}
	return checkSN(node.Next) + 1
}
