class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        root = ListNode()
        node = root
        while True:
            if l1 == None and l2 == None:
                break
            if l1 == None:
                node.next = l2
                break
            if l2 == None:
                node.next = l1
                break
            if l1.val <= l2.val:
                node.next = l1
                l1 = l1.next
            else:
                node.next = l2
                l2 = l2.next
            node = node.next
        return root.next