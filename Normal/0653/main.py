from typing import Dict, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
        
class Solution:
    def findTarget(self, root: Optional[TreeNode], k: int) -> bool:
        hashmap: Dict[int, bool] = {}
        def find(node: Optional[TreeNode], target: int) -> bool:
            if node == None:
                return False
            if hashmap.get(target - node.val) != None:
                return True
            hashmap[node.val] = True
            return find(node.left, target) or find(node.right, target)
        return find(root, k)


        