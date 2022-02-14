from typing import List


class Solution:
    def exchange(self, nums: List[int]) -> List[int]:
        array1: List[int] = []
        array2: List[int] = []
        for value in nums:
            if int(value % 2) != 0:
                array1.append(value)
            else:
                array2.append(value)
        array1.extend(array2)
        return array1
        