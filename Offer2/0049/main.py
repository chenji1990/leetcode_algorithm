import heapq


class Solution:
    def nthUglyNumber(self, n: int) -> int:
        factors = {2, 3, 5}

        h = [1]
        seen = {1: True}
        x, next = 0, 0
        for i in range(1, 1691):
            x = heapq.heappop(h)
            if i == n:
                return x
            for f in factors:
                next = x * f
                if seen.get(next, False) == False:
                    heapq.heappush(h, next)
                    seen[next] = True
        return 1
        
