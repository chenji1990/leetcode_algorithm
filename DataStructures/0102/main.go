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
	array := [][]int{}
	row := []*TreeNode{root}
	nodeArray := []*TreeNode{}
	valueArray := []int{}
	var node *TreeNode
	for len(row) != 0 {
		nodeArray = []*TreeNode{}
		valueArray = []int{}
		for _, node = range row {
			valueArray = append(valueArray, node.Val)
			if node.Left != nil {
				nodeArray = append(nodeArray, node.Left)
			}
			if node.Right != nil {
				nodeArray = append(nodeArray, node.Right)
			}
		}
		array = append(array, valueArray)
		row = nodeArray
	}
	return array
}
