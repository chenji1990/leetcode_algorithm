class Solution {
    func intersect(_ nums1: [Int], _ nums2: [Int]) -> [Int] {
        if nums1.count > nums2.count{
            return intersect(nums2, nums1)
        }

        var hashmap = [Int: Int]()
        for value in nums1{
            if let count = hashmap[value], count > 0{
                hashmap[value] = count + 1
            } else {
                hashmap[value] = 1
            }
        }

        var res = [Int]()
        for value  in nums2{
            if let count = hashmap[value], count > 0{
                res.append(value)
                hashmap[value] = count - 1
            }
        }
        return res
    }
}