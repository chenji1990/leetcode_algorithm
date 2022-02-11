from typing import Dict, List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hashmap: Dict[int, int] = {}
        i1, i2, value = 0, 0, 0

        for i1 in range(len(nums)):
            value = nums[i1]
            i2 = hashmap.get(target - value)
            if i2 != None:
                return [i2, i1]
            else:
                hashmap[value] = i1
        return []