class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None
        
class Solution:
    def deleteNode(self, head: ListNode, val: int) -> ListNode:
        if head == None:
            return None
        if head.val == val:
            return head.next
        
        node1, node2 = head, head.next
        while node2 != None:
            if node2.val == val:
                node1.next = node2.next
                break
            node1 = node1.next
            node2 = node2.next
        return head