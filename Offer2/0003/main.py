from typing import Dict, List


class Solution:
    def findRepeatNumber(self, nums: List[int]) -> int:
        hashmap: Dict[int, bool] = {}
        for v in nums:
            if hashmap.get(v) != None:
                return v
            hashmap.setdefault(v, True)
        return -1