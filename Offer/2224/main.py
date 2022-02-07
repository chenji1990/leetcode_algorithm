class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        root: ListNode = ListNode(0)
        temp: ListNode = None
        while head != None:
            temp = head.next
            head.next = root.next
            root.next = head
            head = temp
        return root.next