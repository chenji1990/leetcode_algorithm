
from typing import List


class Solution:

    def verifyPostorder(self, postorder: List[int]) -> bool:
        def recur(start: int, end: int) -> bool:
            if start >= end: return True
            head = postorder[end]
            leftEnd = start
            while postorder[leftEnd] < head: leftEnd += 1
            rightEnd = leftEnd
            while postorder[rightEnd] > head: rightEnd += 1
            return rightEnd == end and recur(start, leftEnd - 1) and recur(leftEnd, rightEnd - 1)
        return recur(0, len(postorder) - 1)
