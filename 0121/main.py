import math
from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        minPrice: int = 1e9
        price: int = 1e9
        profit: int = 0
        tempProfit: int = 0

        for price in prices:
            if price < minPrice:
                minPrice = price
            else:
                tempProfit = price - minPrice
                if tempProfit > profit:
                    profit = tempProfit
        
        return profit