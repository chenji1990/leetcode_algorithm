from typing import Dict


class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random

class Solution:
    def copyRandomList(self, head: 'Node') -> 'Node':
        if head == None:
            return None

        root = Node(0)
        tempRoot = root
        tempHead = head
        i = 0

        map1: Dict[Node, int] = {}
        map2: Dict[int, Node] = {}

        while tempHead != None:
            tempRoot.next = Node(tempHead.val)
            map1[tempHead] = i
            map2[i] = tempRoot.next

            tempRoot = tempRoot.next
            tempHead = tempHead.next
            i += 1
        
        tempRoot = root
        tempHead = head
        tempRandom: Node | None = None
        while tempHead != None:
            tempRandom = tempHead.random
            if tempRandom != None and (index := map1.get(tempRandom)) != None:
                tempRoot.next.random = map2.get(index)
            tempRoot = tempRoot.next
            tempHead = tempHead.next

        return root.next
        


