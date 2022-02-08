class Solution {
    func merge(_ nums1: inout [Int], _ m: Int, _ nums2: [Int], _ n: Int) {
        var curIndex = m + n - 1
        var p1 = m - 1
        var p2 = n - 1

        var value1 = 0
        var value2 = 0

        while p1 >= 0 || p2 >= 0{
            if p1 < 0{
                nums1[curIndex] = nums2[p2]
                p2 -= 1
                curIndex -= 1
                continue
            }

            if p2 < 0{
                nums1[curIndex] = nums1[p1]
                p1 -= 1
                curIndex -= 1
                continue
            }

            value1 = nums1[p1]
            value2 = nums2[p2]
            if value2 >= value1{
                nums1[curIndex] = value2
                p2 -= 1
            } else {
                nums1[curIndex] = value1
                p1 -= 1
            }
            curIndex -= 1
        }
    }
}

var array = [4,5,6,0,0,0]
let so = Solution()
so.merge(&array, 3, [1,2,3], 3)
print(array)