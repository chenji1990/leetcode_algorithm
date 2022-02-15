package main

import "strings"

func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return false
	}
	length := len(s)
	for i := 0; i < length; i++ {
		if s[i] == 'e' || s[i] == 'E' {
			return (isInt(s[:i], false) || isFloat(s[:i])) && isInt(s[i+1:], false)
		}
	}
	return isInt(s, false) || isFloat(s)
}

func isFloat(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	lastIndex := len(s) - 1
	if lastIndex == -1 {
		return false
	}
	if s[0] == '.' {
		return isInt(s[1:], true)
	}
	if s[lastIndex] == '.' {
		return isInt(s[:lastIndex], false)
	}
	for i := 1; i < lastIndex; i++ {
		if s[i] == '.' {
			return isInt(s[:i], false) && isInt(s[i+1:], true)
		}
	}
	return false
}

func isInt(s string, unsign bool) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '+' || s[0] == '-' {
		if unsign {
			return false
		}
		s = s[1:]
	}
	length := len(s)
	if length == 0 {
		return false
	}
	for i := 0; i < length; i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}
