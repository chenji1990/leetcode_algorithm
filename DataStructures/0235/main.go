package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 在方法一中，我们对从根节点开始，通过遍历找出到达节点 pp 和 qq 的路径，一共需要两次遍历。我们也可以考虑将这两个节点放在一起遍历。
// 整体的遍历过程与方法一中的类似：
// 我们从根节点开始遍历；
// 如果当前节点的值大于 pp 和 qq 的值，说明 pp 和 qq 应该在当前节点的左子树，因此将当前节点移动到它的左子节点；
// 如果当前节点的值小于 pp 和 qq 的值，说明 pp 和 qq 应该在当前节点的右子树，因此将当前节点移动到它的右子节点；
// 如果当前节点的值不满足上述两条要求，那么说明当前节点就是「分岔点」。此时，pp 和 qq 要么在当前节点的不同的子树中，要么其中一个就是当前节点。
// 可以发现，如果我们将这两个节点放在一起遍历，我们就省去了存储路径需要的空间。

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val == q.Val {
		return p
	}
	if p.Val > q.Val {
		p, q = q, p
	}
	for root != nil {
		if root.Val > q.Val {
			root = root.Left
			continue
		}
		if root.Val < p.Val {
			root = root.Right
			continue
		}
		return root
	}
	return root
}
