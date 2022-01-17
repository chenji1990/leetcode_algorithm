from typing import List


class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        p1, p2 = m - 1, n - 1
        value1, value2 = 0, 0

        for i in range(m + n - 1, -1, -1):
            if p1 < 0:
                nums1[i] = nums2[p2]
                p2 -= 1
                continue
            
            if p2 < 0:
                nums1[i] = nums1[p1]
                p1 -= 1
                continue

            value1, value2 = nums1[p1], nums2[p2]
            if value2 >= value1:
                nums1[i] = value2
                p2 -= 1
            else:
                nums1[i] = value1
                p1 -= 1
        
