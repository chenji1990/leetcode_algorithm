package main

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	hashmap := make(map[int]int)
	for _, value := range nums1 {
		if count, ok := hashmap[value]; ok {
			hashmap[value] = count + 1
		} else {
			hashmap[value] = 1
		}
	}
	res := []int{}
	for _, value := range nums2 {
		if count, ok := hashmap[value]; ok && count > 0 {
			res = append(res, value)
			hashmap[value] = count - 1
		}
	}
	return res
}
