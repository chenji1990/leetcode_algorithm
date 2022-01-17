package main

func maxSubArray(nums []int) int {
	preValue := nums[0]
	maxValue := preValue

	for _, value := range nums[1:] {
		if preValue > 0 {
			preValue += value
		} else {
			preValue = value
		}

		if preValue > maxValue {
			maxValue = preValue
		}
	}
	return maxValue
}
