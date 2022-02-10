package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hashmap := make(map[*ListNode]bool)
	for headA != nil {
		hashmap[headA] = true
		headA = headA.Next
	}
	var ok bool
	for headB != nil {
		if _, ok = hashmap[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
