from typing import List, Optional

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def postorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        array: List[int] = []
        def order(node: Optional[TreeNode]):
            if node == None:
                return
            order(node.left)
            order(node.right)
            array.append(node.val)
        order(root)
        return array
        