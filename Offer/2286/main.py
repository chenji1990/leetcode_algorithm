from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        minValue: int = 1e9
        profit: int = 0
        temp: int

        for price in prices:
            temp = price - minValue
            if temp > profit:
                profit = temp
            if price < minValue:
                minValue = price
        
        return profit