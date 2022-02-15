

from typing import List


class Solution:
    def printNumbers(self, n: int) -> List[int]:
        length = 10 ** n
        res : List[int] = [0] * length
        for i in range(1, length):
            res[i] = i
        return res[1:]
