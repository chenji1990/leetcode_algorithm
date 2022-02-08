import collections
from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def levelOrder(self, root: TreeNode) -> List[int]:
        if root == None:
            return []

        array: List[int] = []
        queue = collections.deque([root])
        node: TreeNode
        while len(queue) != 0:
            node = queue.popleft()
            array.append(node.val)
            if node.left != None:
                queue.append(node.left)
            if node.right != None:
                queue.append(node.right)
        return array