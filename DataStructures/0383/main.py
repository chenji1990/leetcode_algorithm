from typing import Dict


class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        if len(ransomNote) > len(magazine):
            return False
        hashmap: Dict[str, int] = {}

        value: str = ""
        for value in magazine:
            hashmap[value] = hashmap.get(value, 0) + 1
        
        for value in ransomNote:
            if hashmap.get(value, 0) > 0:
                hashmap[value] -= 1
            else:
                return False
        
        return True
