package main

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	minValue := numbers[0]
	for _, value := range numbers {
		if value < minValue {
			return value
		}
	}
	return minValue
}
