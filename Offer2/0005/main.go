package main

import "strings"

func replaceSpace(s string) string {
	array := []string{}
	var value rune
	for _, value = range s {
		if value == ' ' {
			array = append(array, "%20")
		} else {
			array = append(array, string(value))
		}
	}
	return strings.Join(array, "")
}
