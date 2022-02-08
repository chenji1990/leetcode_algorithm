from typing import Dict


class Solution:
    def firstUniqChar(self, s: str) -> str:
        if len(s) == 0:
            return " "
        hashmap: Dict[str, int] = {}
        length = len(s)
        for i in range(length):
            if hashmap.get(s[i]) != None:
                hashmap[s[i]] = length
            else:
                hashmap[s[i]] = i
        
        first = sorted(hashmap.items(), key=lambda item:item[1])[0]
        return first[0] if first[1] != length else " "

