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
        tempArray: List[int]
        queue: List[TreeNode] = [root]
        tempQueue: List[TreeNode]
        node: TreeNode
        i, length = 0, 0
        isLToR = True

        while len(queue) != 0:
            tempArray = []
            tempQueue = []
            length = len(queue)
            if isLToR:
                for i in range(length):
                    node = queue[i]
                    tempArray.append(node.val)
                    if node.left != None:
                        tempQueue.append(node.left)
                    if node.right != None:
                        tempQueue.append(node.right)
            else:
                for i in range(length):
                    tempArray.append(queue[length - 1 - i].val)
                    node = queue[i]
                    if node.left != None:
                        tempQueue.append(node.left)
                    if node.right != None:
                        tempQueue.append(node.right)
            res.append(tempArray)
            queue = tempQueue
            isLToR = not isLToR
        return res
        