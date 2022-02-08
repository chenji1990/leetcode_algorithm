class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def isSubStructure(self, A: TreeNode, B: TreeNode) -> bool:
        if A == None or B == None:
            return False
        return self.checkEqual(A, B) or self.isSubStructure(A.left, B) or self.isSubStructure(A.right, B)

    def checkEqual(self, A:TreeNode, B: TreeNode) -> bool:
        if B == None:
            return True
        if A == None or A.val != B.val:
            return False
        return self.checkEqual(A.left, B.left) and self.checkEqual(A.right, B.right)