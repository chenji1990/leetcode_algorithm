class Solution:
    def numWays(self, n: int) -> int:
        if n <= 1:
            return 1
        if n == 2:
            return 2
        
        mod: int = 1e9 + 7
        v1, v2 = 1, 2
        for _ in range(2, n):
            v1, v2 = v2, (v1 + v2) % mod
        return int(v2)