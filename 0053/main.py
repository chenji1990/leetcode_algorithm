from typing import List


class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        preValue = nums[0]
        maxValue = preValue
        
        for value in nums[1:]:
            if preValue > 0:
                preValue += value
            else:
                preValue = value
            if preValue > maxValue:
                maxValue = preValue
        
        return maxValue