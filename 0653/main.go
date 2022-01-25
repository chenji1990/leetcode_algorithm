package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	hashmap := make(map[int]bool)
	ok := false
	var find func(node *TreeNode, target int) bool
	find = func(node *TreeNode, target int) bool {
		if node == nil {
			return false
		}
		if _, ok = hashmap[target-node.Val]; ok {
			return true
		}
		hashmap[node.Val] = true
		return find(node.Left, target) || find(node.Right, target)
	}
	return find(root, k)
}
