from typing import Dict, List


class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        if len(matrix) == 0 or len(matrix[0]) == 0:
            return

        row, column = len(matrix), len(matrix[0])
        i, j, index = 0, 0, 0

        array: List[Dict[int, bool]] = [{}, {}]

        for i in range(row):
            for j in range(column):
                if matrix[i][j] == 0:
                    array[0][i], array[1][j] = True, True

        for i in array[0]:
            for index in range(column):
                matrix[i][index] = 0
        
        for j in array[1]:
            for index in range(row):
                matrix[index][j] = 0

        
        

        