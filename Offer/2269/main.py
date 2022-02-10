from typing import Dict


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> ListNode:
        hashmap: Dict[ListNode, bool] = {}
        while headA != None:
            hashmap[headA] = True
            headA = headA.next
        while headB != None:
            if hashmap.get(headB) != None:
                return headB
            headB = headB.next
        return None