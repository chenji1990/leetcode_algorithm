package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) (res [][]int) {
	var path []int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && sum == node.Val {
			res = append(res, append([]int{}, path...))
			path = path[:len(path)-1]
			return
		}
		dfs(node.Left, sum-node.Val)
		dfs(node.Right, sum-node.Val)
		path = path[:len(path)-1]
	}
	dfs(root, target)
	return res
}
