import collections
from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Codec:

    def serialize(self, root: TreeNode) -> str:
        if root == None: return "[]"
        queue: collections.deque[TreeNode] = collections.deque()
        queue.append(root)
        res: List[str] = []
        while len(queue) != 0:
            node = queue.popleft()
            if node == "None":
                res.append("null")
            else:
                res.append(str(node.val))
                queue.append(node.left)
                queue.append(node.right)
        return "[" + ",".join(res) + "]"

    def deserialize(self, data: str) -> TreeNode:
        if data == "[]": return None
        res: List[str] = data[1:-1].split(",")
        root: TreeNode = TreeNode(int(res[0]))
        queue: collections.deque[TreeNode] = collections.deque()
        queue.append(root)
        i = 1
        while len(queue) != 0:
            node = queue.popleft()
            if res[i] != "null":
                node.left = TreeNode(int(res[i]))
                queue.append(node.left)
            i += 1
            if res[i] != "null":
                node.right = TreeNode(int(res[i]))
                queue.append(node.right)
            i += 1
        return root