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

	var res [][]int
	var tempArray []int
	var queue = []*TreeNode{root}
	var tempQueue []*TreeNode
	var node *TreeNode

	var isLToR = true
	var i, length int

	for len(queue) != 0 {
		tempArray = []int{}
		tempQueue = []*TreeNode{}
		length = len(queue) - 1
		if isLToR {
			for i = 0; i < length; i++ {
				node = queue[i]
				tempArray = append(tempArray, node.Val)
				if node.Left != nil {
					tempQueue = append(tempQueue, node.Left)
				}
				if node.Right != nil {
					tempQueue = append(tempQueue, node.Right)
				}
			}
		} else {
			for i = 0; i < length; i++ {
				tempArray = append(tempArray, queue[length-1-i].Val)
				node = queue[i]
				if node.Left != nil {
					tempQueue = append(tempQueue, node.Left)
				}
				if node.Right != nil {
					tempQueue = append(tempQueue, node.Right)
				}
			}
		}
		res = append(res, tempArray)
		queue = tempQueue
		isLToR = !isLToR
	}
	return res
}
