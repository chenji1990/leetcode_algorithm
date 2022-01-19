package main

func firstUniqChar(s string) int {
	hashmap := make(map[rune]int)

	var index int
	var value rune
	var ok bool
	for index, value = range s {
		if _, ok = hashmap[value]; ok {
			hashmap[value] = -1
		} else {
			hashmap[value] = index
		}
	}

	var minIndex int = 1e9
	for value, index = range hashmap {
		if index != -1 && index < minIndex {
			minIndex = index
		}
	}
	if minIndex == 1e9 {
		return -1
	}
	return minIndex
}
