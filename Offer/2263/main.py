class Solution:
    def translateNum(self, num: int) -> int:
        if num < 10:
            return 1
        src = str(num)
        a, b, res = 0, 1, 1
        temp = ""
        for i in range(1, len(src)):
            a, b = b, res
            temp = src[i - 1 : i + 1]
            if temp >= "10" and temp <= "25":
                res += a
        return res