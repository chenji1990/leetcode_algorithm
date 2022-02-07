from typing import List


class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        mid, left, right = 0, 0, len(nums) - 1
        while left <= right:
            mid = (left + right) // 2
            if mid == nums[mid]:
                left = mid + 1
            else:
                right = mid - 1
        return left