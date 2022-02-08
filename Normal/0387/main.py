from typing import Dict


class Solution:
    def firstUniqChar(self, s: str) -> int:
        hashmap: Dict[str, int] = {}
        index: int = 0
        value: str = ""
        for index, value in enumerate(s):
            if hashmap.get(value) != None:
                hashmap[value] = -1
            else:
                hashmap[value] = index

        minIndex: int = 1e9
        for (value, index) in hashmap.items():
            if index != -1 and index < minIndex:
                minIndex = index
        
        if minIndex == 1e9:
            return -1
        return minIndex