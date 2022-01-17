function maxSubArray(nums: number[]): number {
    let preValue = nums[0]
    let maxValue = preValue

    for (let value of nums.slice(1)) {
        if (preValue > 0) {
            preValue += value
        } else {
            preValue = value
        }
        if (preValue > maxValue) {
            maxValue = preValue
        }
    }
    return maxValue
};