class Solution:
    def countDigitOne(self, n: int) -> int:
        value = 1
        temp: int = 0
        res: int = 0
        while value <= n:
            temp = value * 10
            res += (n // temp) * value + min(max(n % temp - value + 1, 0), value)
            value = temp
        return res