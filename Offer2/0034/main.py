from typing import List


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def pathSum(self, root: TreeNode, target: int) -> List[List[int]]:
        res: List[List[int]] = []
        path: List[int] = []

        def dfs(node: TreeNode, sum: int):
            if node == None:
                return
            nonlocal path
            path.append(node.val)
            if node.left == None and node.right == None and sum == node.val:
                res.append(path.copy())
                path = path[:len(path) - 1]
                return
            dfs(node.left, sum - node.val)
            dfs(node.right, sum - node.val)
            path = path[:len(path) - 1]

        dfs(root, target)
        return res
