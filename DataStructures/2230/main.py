from collections import deque


class MinStack:
    stack: deque[int] 
    minStack: deque[int] 

    def __init__(self):
        self.stack = deque([])
        self.minStack = deque([])


    def push(self, x: int) -> None:
        if len(self.stack) == 0:
            self.minStack.append(x)
        else:
            self.minStack.append(min(x, self.minStack[-1]))
        self.stack.append(x)

    def pop(self) -> None:
        if len(self.stack) != 0:
            self.stack.pop()
            self.minStack.pop()

    def top(self) -> int:
        if len(self.stack) == 0:
            return 0
        return self.stack[-1]

    def min(self) -> int:
        if len(self.minStack) == 0:
            return 0
        return self.minStack[-1]