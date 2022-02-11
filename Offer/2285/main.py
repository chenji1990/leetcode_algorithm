from typing import Dict, List


class Solution:
    def isStraight(self, nums: List[int]) -> bool:
        hashmap: Dict[int, bool] = {}
        minValue, maxValue = 13, 1
        for num in nums:
            if hashmap.get(num) != None:
                return False
            if num != 0:
                hashmap[num] = True
                if num < minValue:
                    minValue = num
                if num > maxValue:
                    maxValue = num
        return maxValue - minValue < 5