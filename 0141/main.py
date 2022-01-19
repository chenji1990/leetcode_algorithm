from typing import Dict, Optional


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        hashmap: Dict[ListNode, bool] = {}
        while head != None:
            if hashmap.get(head) == True:
                return True
            hashmap[head] = True
            head = head.next
        return False