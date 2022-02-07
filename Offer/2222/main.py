from typing import List


class Solution:
    def replaceSpace(self, s: str) -> str:
        array: List[str] = []
        value: str = ""
        for value in s:
            if value == " ":
                array.append("%20")
            else:
                array.append(value)
        return "".join(array)