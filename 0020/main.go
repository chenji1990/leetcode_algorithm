package main

func isValid(s string) bool {
	length := len(s)
	if length <= 0 || length%2 != 0 {
		return false
	}

	hashmap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}
	var lastIndex, i int

	for i = 0; i < length; i++ {
		lastIndex = len(stack) - 1
		if lastIndex <= -1 || stack[lastIndex] != hashmap[s[i]] {
			stack = append(stack, s[i])
			continue
		}
		stack = stack[:lastIndex]
	}
	return len(stack) == 0
}
