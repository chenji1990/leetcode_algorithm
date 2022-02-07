from collections import deque
from typing import List


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def reversePrint(self, head: ListNode) -> List[int]:
        stack: deque[int] = deque([])
        while head != None:
            stack.append(head.val)
            head = head.next
        res: List[int] = []
        while len(stack) > 0:
            res.append(stack.pop())
        return res