class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def kthLargest(self, root: TreeNode, k: int) -> int:
        res: int = 0
        def search(node: TreeNode):
            nonlocal res, k
            if node == None or k <= 0:
                return 
            search(node.right)
            k -= 1
            if k == 0:
                res = node.val
                return
            search(node.left)
            
        search(root)
        return res