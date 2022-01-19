package main

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	hashmap := make(map[rune]int)
	var value rune
	for _, value = range magazine {
		hashmap[value] += 1
	}

	for _, value = range ransomNote {
		if hashmap[value] > 0 {
			hashmap[value] -= 1
		} else {
			return false
		}
	}

	return true
}
