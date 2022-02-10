package main

import "strings"

func reverseWords(s string) string {
	i, j := len(s)-2, len(s)-1
	var array []string
	for i >= -1 {
		if s[j] == ' ' {
			j--
			i--
			continue
		}
		if i == -1 || s[i] == ' ' {
			array = append(array, s[i+1:j+1])
			j = i - 1
			i -= 2
		} else {
			i--
		}
	}
	return strings.Join(array, " ")
}

// "the sky is blue"
