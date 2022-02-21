package main

func countDigitOne(n int) (res int) {
	value := 1
	var temp int
	for value <= n {
		temp = value * 10
		res += (n/temp)*value + min(max(n%temp-value+1, 0), value)
		value = temp
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。

// 例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。
