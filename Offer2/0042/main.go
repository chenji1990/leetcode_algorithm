package main

func maxSubArray(nums []int) int {
	maxValue, length := nums[0], len(nums)
	for i := 1; i < length; i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > maxValue {
			maxValue = nums[i]
		}
	}
	return maxValue
}
