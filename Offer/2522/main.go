package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	array := []int{}
	if root == nil {
		return array
	}

	queue := []*TreeNode{root}
	var node *TreeNode
	for len(queue) != 0 {
		node = queue[0]
		queue = queue[1:]

		array = append(array, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return array
}
