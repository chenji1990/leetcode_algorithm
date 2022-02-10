from typing import List


class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        if len(board) == 0 or len(board[0]) == 0:
            return False

        row, column, length = len(board), len(board[0]), len(word)
        if row * column < length:
            return False

        def dfs(i: int, j: int, k: int) -> bool:
            if k >= length:
                return True
            if i >= 0 and i < row and j >= 0 and j < column and board[i][j] == word[k]:
                board[i][j] = "0"
                ok = dfs(i-1, j, k+1) or dfs(i+1, j, k+1) or dfs(i, j-1, k+1) or dfs(i, j+1, k+1)
                board[i][j] = word[k]
                return ok
            return False

        i, j = 0, 0
        for i in range(row):
            for j in range(column):
                if dfs(i, j, 0):
                    return True
        return False
