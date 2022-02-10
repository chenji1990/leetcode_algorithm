from typing import List


class Solution:
    def maxValue(self, grid: List[List[int]]) -> int:
        if len(grid) == 0 or len(grid[0]) == 0:
            return 0
        row, column = len(grid), len(grid[0])
        i, j = 0, 0
        for i in range(1, row):
            grid[i][0] += grid[i - 1][0]
        for j in range(1, column):
            grid[0][j] += grid[0][j - 1]
        for i in range(1, row):
            for j in range(1, column):
                grid[i][j] += max(grid[i - 1][j], grid[i][j - 1])
        return grid[-1][-1]
