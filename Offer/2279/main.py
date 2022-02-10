from typing import List


class Solution:
    def reverseWords(self, s: str) -> str:
        i, j = len(s) - 2, len(s) - 1
        array: List[str] = []
        while i >= -1:
            if s[j] == " ":
                i -= 1
                j -= 1
                continue

            if i == -1 or s[i] == " ":
                array.append(s[i + 1: j + 1])
                j = i - 1
                i -= 2
            else:
                i -= 1
        return " ".join(array)
