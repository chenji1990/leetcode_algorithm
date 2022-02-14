from typing import List


class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        maxValue = nums[0]
        for i in range(1, len(nums)):
            if nums[i - 1] > 0:
                nums[i] += nums[i - 1]
            if nums[i] > maxValue:
                maxValue = nums[i]
        return maxValue