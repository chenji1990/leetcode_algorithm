from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        count: int = 0
        for v in nums:
            if target == v:
                count += 1
            elif target < v:
                return count
        return count