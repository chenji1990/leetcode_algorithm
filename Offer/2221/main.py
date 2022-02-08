from typing import List


class Solution:
    def findNumberIn2DArray(self, matrix: List[List[int]], target: int) -> bool:
        if len(matrix) == 0 or len(matrix[0]) == 0:
            return False
        row, column = len(matrix), len(matrix[0])
        i, j, value = 0, 0, 0
        for i in range(row):
            for j in range(column):
                value = matrix[i][j] - target
                if value == 0:
                    return True
                elif value < 0:
                    continue
                else:
                    break
        return False
