from typing import List


class Solution:
    def movingCount(self, m: int, n: int, k: int) -> int:
        matrix: List[List[bool]] = [
            [False for _ in range(n)] for _ in range(m)]

        def dfs(i: int, j: int) -> int:
            if i < 0 or i >= m or j < 0 or j >= n or (self.addBit(i) + self.addBit(j) > k) or matrix[i][j]:
                return 0
            matrix[i][j] = True
            return dfs(i, j + 1) + dfs(i, j - 1) + dfs(i + 1, j) + dfs(i - 1, j) + 1

        return dfs(0, 0)

    def addBit(self, v) -> int:
        sum = 0
        for i in str(v):
            sum += int(i)
        return sum
