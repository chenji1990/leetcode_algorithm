class Solution:
    def isNumber(self, s: str) -> bool:
        s = s.strip()
        length = len(s)
        if length == 0:
            return False
        for i in range(length):
            if s[i] == "e" or s[i] == "E":
                return (self.isInt(s[:i]) or self.isFloat(s[:i])) and self.isInt(s[i+1:])
        return self.isInt(s) or self.isFloat(s)

    def isFloat(self, s: str) -> bool:
        if len(s) == 0:
            return False
        if s[0] == "+" or s[0] == "-":
            s = s[1:]
        lastIndex = len(s) - 1
        if lastIndex == -1:
            return False
        if s[0] == ".":
            return self.isInt(s[1:], True)
        if s[lastIndex] == ".":
            return self.isInt(s[:lastIndex])
        for i in range(1, lastIndex):
            if s[i] == ".":
                return (self.isInt(s[:i]) or self.isFloat(s[:i])) and self.isInt(s[i + 1:])
        return False

    def isInt(self, s: str, unsign: bool = False) -> bool:
        if len(s) == 0:
            return False
        if s[0] == "+" or s[0] == "-":
            if unsign:
                return False
            s = s[1:]

        length = len(s)
        if length == 0:
            return False
        for i in range(length):
            if s[i] < "0" or s[i] > "9":
                return False
        return True
