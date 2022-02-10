from typing import Dict


class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        length = len(s)
        if length <= 1:
            return length
        hashmap: Dict[str,int] = {}
        i, j, res, temp = 0, 0, 0, 0
        for j in range(length):
            if (i := hashmap.get(s[j])) != None and temp >= j - i:
                temp = j - i
            else:
                temp += 1
            if temp > res:
                res = temp
            hashmap[s[j]] = j
        return res