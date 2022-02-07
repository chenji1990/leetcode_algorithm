package main

func findRepeatNumber(nums []int) int {
	hashmap := make(map[int]bool)
	ok := false
	for _, v := range nums {
		if _, ok = hashmap[v]; ok {
			return v
		}
		hashmap[v] = true
	}
	return -1
}
