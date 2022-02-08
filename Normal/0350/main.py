from typing import Dict, List


# [4,9,5]
# [9,4,9,8,4]

class Solution:
    def intersect(self, nums1: List[int], nums2: List[int]) -> List[int]:
        if len(nums1) > len(nums2):
            return self.intersect(nums2, nums1)
        
        hashmap: Dict[int,int] = {}
        value: int = None
        count: int = None

        for value in nums1:
            if (count := hashmap.get(value)) != None:
                hashmap[value] = count + 1
            else:
                hashmap[value] = 1

        res: List[int] = []
        for value in nums2:
            if (count := hashmap.get(value)) != None and count > 0:
                res.append(value)
                hashmap[value] = count - 1       
        return res
