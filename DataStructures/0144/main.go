package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	array := []int{}
	var order func(node *TreeNode)
	order = func(node *TreeNode) {
		if node == nil {
			return
		}
		array = append(array, node.Val)
		order(node.Left)
		order(node.Right)
	}
	order(root)
	return array
}
