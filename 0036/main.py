from typing import List


class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        array: List[List[List[int]]] = []
        for _ in range(9):
            array.append([
                [0, 0, 0, 0, 0, 0, 0, 0, 0],
                [0, 0, 0, 0, 0, 0, 0, 0, 0],
                [0, 0, 0, 0, 0, 0, 0, 0, 0],
            ])
        
        i, j, index, section = 0, 0, 0, 0

        for i in range(9):
            for j in range(9):
                if board[i][j] == ".":
                    continue

                index = ord(board[i][j]) - 49
                if array[index][0][i] > 0 or array[index][1][j] > 0:
                    return False

                section = int(i / 3) * 3 + int(j / 3)
                if array[index][2][section] > 0:
                    return False

                array[index][0][i] = 1
                array[index][1][j] = 1
                array[index][2][section] = 1
        
        return True

obj = Solution()
res = obj.isValidSudoku([[".",".","4",".",".",".","6","3","."],[".",".",".",".",".",".",".",".","."],["5",".",".",".",".",".",".","9","."],[".",".",".","5","6",".",".",".","."],["4",".","3",".",".",".",".",".","1"],[".",".",".","7",".",".",".",".","."],[".",".",".","5",".",".",".",".","."],[".",".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".",".","."]])

print(res)