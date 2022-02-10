package main

func exchange(nums []int) []int {
	var array1, array2 []int
	for _, value := range nums {
		if value%2 != 0 {
			array1 = append(array1, value)
		} else {
			array2 = append(array2, value)
		}
	}
	return append(array1, array2...)
}
