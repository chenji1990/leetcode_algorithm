package main

func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	var temp int
	for i < j {
		temp = nums[i] + nums[j]
		if temp < target {
			i++
			continue
		}
		if temp > target {
			j--
			continue
		}
		return []int{nums[i], nums[j]}
	}
	return []int{}
}
