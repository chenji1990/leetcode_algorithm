class Solution:
    def findNthDigit(self, n: int) -> int:
        digit, digitNum, count = 1, 1, 9
        while n > count:
            n -= count
            digit += 1
            digitNum *= 10
            count = 9 * digit * digitNum
        num = digitNum + (n - 1) // digit
        index = (n - 1) % digit
        numStr = str(num)
        return int(numStr[index]) - int("0")

