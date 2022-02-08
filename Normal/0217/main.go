package main

func containsDuplicate(nums []int) bool {
	hashmap := make(map[int]bool)
	for _, value := range nums {
		if _, ok := hashmap[value]; ok {
			return true
		}
		hashmap[value] = true
	}
	return false
}
