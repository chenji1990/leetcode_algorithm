from typing import Dict, List

class Solution:
    def isValid(self, s: str) -> bool:
        length = len(s)
        if length == 0 or length % 2 != 0:
            return False

        hashmap: Dict[str, str] = {
            ")": "(",
            "]": "[",
            "}": "{",
        }

        stack: List[str] = []
        lastIndex: int = 0
        i: int = 0

        for i in range(length):
            lastIndex = len(stack) - 1
            if lastIndex <= -1 or stack[lastIndex] != hashmap.get(s[i], ""):
                stack.append(s[i])
                continue
            stack = stack[:lastIndex]
        return len(stack) == 0