package main

func isStraight(nums []int) bool {
	hashmap := make(map[int]bool)
	var ok bool
	minValue, maxValue := 13, 1
	for _, num := range nums {
		if _, ok = hashmap[num]; ok {
			return false
		}
		if num != 0 {
			hashmap[num] = true
			if num < minValue {
				minValue = num
			}
			if num > maxValue {
				maxValue = num
			}
		}
	}
	return maxValue-minValue < 5
}
