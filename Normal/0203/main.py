class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def removeElements(self, head: ListNode, val: int) -> ListNode:
        res = ListNode()
        header = res
        while head != None:
            if head.val != val:
                res.next = head
                res = res.next
                head = head.next
                continue
            head = head.next
        res.next = None
        return header.next