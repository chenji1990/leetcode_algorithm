from typing import List


class Solution:
    def validateStackSequences(self, pushed: List[int], popped: List[int]) -> bool:
        stack: List[int] = []
        i = 0
        for value in pushed:
            stack.append(value)
            while len(stack) > 0 and stack[-1] == popped[i]:
                stack = stack[:len(stack) - 1]
                i += 1
        return len(stack) == 0