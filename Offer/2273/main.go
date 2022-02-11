package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) (res int) {
	var search func(*TreeNode)
	search = func(node *TreeNode) {
		if node == nil || k <= 0 {
			return
		}
		search(node.Right)
		k--
		if k == 0 {
			res = node.Val
			return
		}
		search(node.Left)
	}
	search(root)
	return
}

// 给定一棵二叉搜索树，请找出其中第 k 大的节点的值。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthLargest2(root *TreeNode, k int) (res int) {
	var search func(root *TreeNode)
	search = func(node *TreeNode) {
		if node == nil || k <= 0 {
			return
		}
		search(node.Right)
		k--
		if k == 0 {
			res = node.Val
			return
		}
		search(node.Left)
	}
	search(root)
	return
}
