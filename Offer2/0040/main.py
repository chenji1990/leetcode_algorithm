from typing import List


class Solution:
    def getLeastNumbers(self, arr: List[int], k: int) -> List[int]:
        count = min(len(arr), k)
        arr.sort()
        return arr[:count]

