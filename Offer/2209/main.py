from collections import deque


class CQueue:
    stack1: deque[int] 
    stack2: deque[int] 
    def __init__(self):
        self.stack1 = deque([])
        self.stack2 = deque([])

    def appendTail(self, value: int) -> None:
        self.stack1.append(value)

    def deleteHead(self) -> int:
        if len(self.stack2) == 0:
            while len(self.stack1) > 0:
                self.stack2.append(self.stack1.pop())
                
        if len(self.stack2) > 0:
            return self.stack2.pop()
        return -1