package main

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length <= 1 {
		return length
	}
	hashmap := make(map[byte]int)
	i, j, res, temp := 0, 0, 0, 0
	var ok bool
	for j = 0; j < length; j++ {
		if i, ok = hashmap[s[j]]; ok && temp >= j-i {
			temp = j - i
		} else {
			temp += 1
		}
		if res < temp {
			res = temp
		}
		hashmap[s[j]] = j
	}
	return res
}
