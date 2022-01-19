package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	hashmap := make(map[*ListNode]bool)
	for head != nil {
		if hashmap[head] {
			return true
		}
		hashmap[head] = true
		head = head.Next
	}
	return false
}
