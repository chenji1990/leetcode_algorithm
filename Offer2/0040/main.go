package main

import "sort"

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	count := min(len(arr), k)
	return arr[:count]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
