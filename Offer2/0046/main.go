package main

import "strconv"

func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	src := strconv.Itoa(num)
	length := len(src)

	a, b, res := 0, 1, 1 // 初始种类
	temp := ""
	for i := 1; i < length; i++ {
		a, b = b, res
		temp = src[i-1 : i+1]
		if temp >= "10" && temp <= "25" {
			res += a
		}
	}
	return res
}
