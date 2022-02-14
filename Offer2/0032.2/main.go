package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	tempArray := []int{}
	queue := []*TreeNode{root}
	tempQueue := []*TreeNode{}
	var node *TreeNode
	for len(queue) != 0 {
		tempArray = []int{}
		tempQueue = []*TreeNode{}
		for _, node = range queue {
			tempArray = append(tempArray, node.Val)
			if node.Left != nil {
				tempQueue = append(tempQueue, node.Left)
			}
			if node.Right != nil {
				tempQueue = append(tempQueue, node.Right)
			}
		}
		res = append(res, tempArray)
		queue = tempQueue
	}

	return res
}
