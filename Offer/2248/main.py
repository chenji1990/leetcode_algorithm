import collections
from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        if root == None:
            return []

        res: List[List[int]] = []
        tempArray: List[int] = []
        queue: List[TreeNode] = [root]
        tempQueue: List[TreeNode] = []
        node: TreeNode

        while len(queue) != 0:
            tempArray = []
            tempQueue = []
            for node in queue:
                tempArray.append(node.val)
                if node.left != None:
                    tempQueue.append(node.left)
                if node.right != None:
                    tempQueue.append(node.right)
            res.append(tempArray)
            queue = tempQueue

        return res
