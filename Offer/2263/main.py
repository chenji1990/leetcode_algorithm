from typing import List


class Solution:
    def minNumber(self, nums: List[int]) -> str:
        array: List[str] = [str(x) for x in nums]
        for i in range(len(array) - 1):
            for j in range(i + 1, len(array)):
                if array[i] + array[j] > array[j] + array[i]:
                    array[i], array[j] = array[j], array[i]
        return "".join(array)
