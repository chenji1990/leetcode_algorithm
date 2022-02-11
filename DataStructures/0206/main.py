class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        res = ListNode()
        temp: ListNode = None
        while head != None:
            temp = head.next
            head.next = res.next
            res.next = head
            head = temp
        return res.next
        