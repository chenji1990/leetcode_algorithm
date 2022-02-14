from typing import Tuple


class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def treeToDoublyList(self, root: 'Node') -> 'Node':
        if root == None:
            return None

        temp1, temp2 = None, None

        def dfs(node) -> Tuple:
            if node.left == None and node.right == None:
                return node, node

            nonlocal temp1, temp2

            if node.right == None:
                temp1 = dfs(node.left)
                node.left, temp1[1].right = temp1[1], node
                return temp1[0], node

            if node.left == None:
                temp2 = dfs(node.right)
                node.right, temp2[0].left = temp2[0], node
                return node, temp2[1]

            temp1, temp2 = dfs(node.left), dfs(node.right)
            node.left, temp1[1].right = temp1[1], node
            node.right, temp2[0].left = temp2[0], node
            return temp1[0], temp2[1]

        temp1 = dfs(root)
        temp1[0].left, temp1[1].right = temp1[1], temp1[0]
        return temp1[0]
