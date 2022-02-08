class Solution:
    def fib(self, n: int) -> int:
        if n <= 1:
            return n
        mod: int = 1e9 + 7
        v1: int = 0
        v2: int = 1
        for _ in range(1, n):
            v1, v2 = v2, (v1 + v2) % mod
        return int(v2)