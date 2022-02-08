from typing import List


class Solution:
    def matrixReshape(self, mat: List[List[int]], r: int, c: int) -> List[List[int]]:
        if len(mat) == 0 or len(mat[0]) == 0 or r * c != len(mat) * len(mat[0]):
            return mat
        
        res: List[List[int]] = []
        newRowArray: List[int] = []
        column = 0

        for rowArray in mat:
            for value in rowArray:
                if column < c:
                    newRowArray.append(value)
                    column += 1
                else:
                    res.append(newRowArray)
                    newRowArray = [value]
                    column = 1
        res.append(newRowArray)
        return res