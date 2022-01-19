class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        res = ListNode(-1e9)
        header = res
        while head != None:
            if head.val != res.val:
                res.next = head
                head = head.next
                res = res.next
                continue
            head = head.next
        res.next = None
        return header.next
