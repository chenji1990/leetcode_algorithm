package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	array := []int{}
	var order func(node *TreeNode)
	order = func(node *TreeNode) {
		if node == nil {
			return
		}
		order(node.Left)
		array = append(array, node.Val)
		order(node.Right)
	}
	order(root)
	return array
}
