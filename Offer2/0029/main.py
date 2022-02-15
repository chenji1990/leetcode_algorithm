from typing import List


class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        if len(matrix) == 0 or len(matrix[0]) == 0:
            return []
        row, column = len(matrix), len(matrix[0])
        if row == 1:
            return matrix[0]
        if column == 1:
            return list(matrix[i][0] for i in range(row))
        res: List[int] = []
        subMatrix: List[List[int]] = []

        i, j = 0, 0
        for j in range(column):
            res.append(matrix[i][j])

        j = column - 1
        for i in range(1, row - 1):
            res.append(matrix[i][j])
            subMatrix.append(matrix[i][1:j])

        i = row - 1
        for j in range(column - 1, -1, -1):
            res.append(matrix[i][j])

        j = 0
        for i in range(row - 2, 0, -1):
            res.append(matrix[i][j])
        
        res.extend(self.spiralOrder(subMatrix))
        return res
