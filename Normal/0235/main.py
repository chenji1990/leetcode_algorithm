class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def lowestCommonAncestor(self, root: TreeNode, p: TreeNode, q: TreeNode) -> TreeNode:
        if p.val > q.val:
            temp = p
            p = q
            q = temp
        while root != None:
            if root.val > q.val:
                root = root.left
                continue
            if root.val < p.val:
                root = root.right
                continue
            return root
        return root