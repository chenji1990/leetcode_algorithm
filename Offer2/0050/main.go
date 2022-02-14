package main

func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
	length := len(s)
	hashmap := make(map[byte]int)
	var i int
	var value byte
	var ok bool

	for i = 0; i < length; i++ {
		value = s[i]
		if _, ok = hashmap[value]; ok {
			hashmap[value] = length
		} else {
			hashmap[value] = i
		}
	}

	var minK byte = ' '
	var minV int = length
	for k, v := range hashmap {
		if v < minV {
			minK = k
			minV = v
		}
	}
	return minK
}
