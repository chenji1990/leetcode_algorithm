from typing import Dict


class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        hashmap: Dict[str, int] = {}
        value = ""
        count = 0
        for value in s:
            hashmap[value] = hashmap.get(value, 0) + 1
        for value in t:
            count = hashmap.get(value, 0)
            if count > 0:
                hashmap[value] = count - 1
            else:
                return False

        return True
