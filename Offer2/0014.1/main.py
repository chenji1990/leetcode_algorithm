class Solution:
    def cuttingRope(self, n: int) -> int:
        array = [0] * (n + 1)
        array[0], array[1], array[2] = 0, 1, 1
        j, half = 0, 0
        for i in range(3, n+1):
            array[i] = array[i - 1]
            half = i // 2
            for j in range(1, half + 1):
                array[i] = max(array[i], max(array[j], j) * max(array[i - j], i - j))
        return array[n]