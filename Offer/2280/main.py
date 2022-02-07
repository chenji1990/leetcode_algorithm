from typing import List


class Solution:
    def reverseLeftWords(self, s: str, n: int) -> str:
        if n <= 0:
            return s
        
        leftArray: List[str] = []
        rightArray: List[str] = []

        for i in range(0, len(s)):
            if i < n:
                rightArray.append(s[i])
            else:
                leftArray.append(s[i])

        leftArray.extend(rightArray)
        return "".join(leftArray)