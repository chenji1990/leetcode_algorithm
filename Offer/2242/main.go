package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return checkEqual(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func checkEqual(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}

	if A == nil || A.Val != B.Val {
		return false
	}

	return checkEqual(A.Left, B.Left) && checkEqual(A.Right, B.Right)
}
