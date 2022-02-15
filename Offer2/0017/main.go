package main

import (
	"math"
)

func printNumbers(n int) []int {
	length := int(math.Pow10(n))
	res := make([]int, length)
	for i := 1; i < length; i++ {
		res[i] = i
	}
	return res[1:]
}
