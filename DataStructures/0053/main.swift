class Solution {
    func maxSubArray(_ nums: [Int]) -> Int {
        var preValue = nums[0]
        var maxValue = preValue
        for value in nums.dropFirst(){
            if preValue > 0{
                preValue += value
            } else {
                preValue = value
            }
            if preValue > maxValue{
                maxValue = preValue
            }
        }
        return maxValue
    }
}