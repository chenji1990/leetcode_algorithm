from typing import List


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
        
class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        if root == None:
            return []
        array: List[List[int]] = []
        row: List[TreeNode] = [root]
        nodeArray: List[TreeNode] = []
        valueArray: List[int] = []
        node: TreeNode = None
        while len(row) != 0:
            nodeArray = []
            valueArray = []
            for node in row:
                valueArray.append(node.val)
                if node.left != None:
                    nodeArray.append(node.left)
                if node.right != None:
                    nodeArray.append(node.right)
            row = nodeArray
            array.append(valueArray)
        return array