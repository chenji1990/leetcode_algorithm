package main

func search(nums []int, target int) int {
	count := 0
	for _, v := range nums {
		if v == target {
			count += 1
		} else if v > target {
			return count
		}
	}
	return count
}
