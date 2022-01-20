from typing import List


class MyQueue:
    stack: List[int]
    def __init__(self):
        self.stack = []

    def push(self, x: int) -> None:
        self.stack.append(x)

    def pop(self) -> int:
        length = len(self.stack)
        if length <= 0:
            return 0
        firstValue = self.stack[0]
        self.stack = self.stack[1: length]
        return firstValue

    def peek(self) -> int:
        if len(self.stack) == 0:
            return 0
        return self.stack[0]

    def empty(self) -> bool:
        return len(self.stack) == 0



# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()