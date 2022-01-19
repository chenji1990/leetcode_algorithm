package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashmap := make(map[rune]int)
	var value rune
	for _, value = range s {
		hashmap[value] += 1
	}
	for _, value = range t {
		if hashmap[value] > 0 {
			hashmap[value] -= 1
		} else {
			return false
		}
	}
	return true
}
