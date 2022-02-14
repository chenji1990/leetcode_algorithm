class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def getKthFromEnd(self, head: ListNode, k: int) -> ListNode:
        length = self.checkSN(head)
        if length < k:
            return None
        for _ in range(length - k):
            head = head.next
        return head

    def checkSN(self, node: ListNode) -> int:
        if node == None:
            return 0
        if node.next == None:
            return 1
        return self.checkSN(node.next) + 1