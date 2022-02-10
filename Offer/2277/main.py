from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        i, j = 0, len(nums) - 1
        temp: int = 0
        while i < j:
            temp = nums[i] + nums[j]
            if temp < target:
                i += 1
                continue
            if temp > target:
                j -= 1
                continue
            return [nums[i], nums[j]]
        return []
                