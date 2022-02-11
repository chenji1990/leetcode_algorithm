from typing import List


class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        if numRows <= 0:
            return []
        if numRows == 1:
            return [[1]]
        if numRows == 2:
            return [[1], [1, 1]]

        res: List[List[int]] = [[1], [1, 1]]
        rowArray: List[int] = [1]
        row, column = 0, 0

        for row in range(2, numRows):
            for column in range(1, row):
                rowArray.append(res[row - 1][column - 1] + res[row - 1][column])
            rowArray.append(1)
            res.append(rowArray)
            rowArray = [1]
        
        return res