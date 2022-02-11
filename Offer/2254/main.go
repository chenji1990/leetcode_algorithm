package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}

	var l1, l2, r1, r2 *Node

	var dfs func(*Node) (*Node, *Node)
	dfs = func(node *Node) (*Node, *Node) {
		if node == nil {
			return nil, nil
		}
		if node.Left == nil && node.Right == nil {
			return node, node
		}
		if node.Right == nil {
			l1, l2 = dfs(node.Left)
			node.Left, l2.Right = l2, node
			return l1, node
		}
		if node.Left == nil {
			r1, r2 = dfs(node.Right)
			node.Right, r1.Left = r1, node
			return node, r2
		}
		l1, l2 = dfs(node.Left)
		node.Left, l2.Right = l2, node
		r1, r2 = dfs(node.Right)
		node.Right, r1.Left = r1, node
		return l1, r2
	}
	l1, r2 = dfs(root)
	l1.Left, r2.Right = r2, l1
	return l1
}
