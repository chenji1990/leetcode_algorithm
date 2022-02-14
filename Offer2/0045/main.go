package main

import (
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	array := []string{}
	for _, num := range nums {
		array = append(array, strconv.Itoa(num))
	}
	length := len(array)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if array[i]+array[j] > array[j]+array[i] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return strings.Join(array, "")
}
