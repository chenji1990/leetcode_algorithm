class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def isValidBST(self, root: TreeNode) -> bool:
        return self.isValid(root, -1e32, 1e31-1)

    def isValid(self, node: TreeNode, lower: int, upper: int) -> bool:
        if node == None:
            return True
        if node.val <= lower or node.val >= upper:
            return False
        return self.isValid(node.left, lower, node.val) and self.isValid(node.right, node.val, upper)