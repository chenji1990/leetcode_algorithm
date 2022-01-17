package main

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	i1, i2, value, ok := 0, 0, 0, false
	for i1, value = range nums {
		if i2, ok = hashmap[target-value]; ok {
			return []int{i2, i1}
		}
		hashmap[value] = i1
	}
	return []int{}
}
