class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def isBalanced(self, root: TreeNode) -> bool:
        if root == None:
            return True
        if abs(self.checkDepth(root.left) - self.checkDepth(root.right)) > 1:
            return False
        return self.isBalanced(root.left) and self.isBalanced(root.right)

    def checkDepth(self, node: TreeNode) -> int:
        if node == None:
            return 0
        return max(self.checkDepth(node.left), self.checkDepth(node.right)) + 1