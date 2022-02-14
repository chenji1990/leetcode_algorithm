from typing import List


class Solution:
    def minArray(self, numbers: List[int]) -> int:
        if len(numbers) == 0:
            return -1
        minValue = numbers[0]
        for value in numbers:
            if value < minValue:
                return value
        return minValue