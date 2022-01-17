class Solution {
    func twoSum(_ nums: [Int], _ target: Int) -> [Int] {
        var hashmap = [Int: Int]()
        for (index1, value) in nums.enumerated(){
            if let index2 = hashmap[target - value]{
                return [index2, index1]
            }
            hashmap[value] = index1
        }
        return []
    }
}