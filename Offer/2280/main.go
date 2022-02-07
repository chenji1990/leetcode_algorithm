package main

func reverseLeftWords(s string, n int) string {
	if n <= 0 {
		return s
	}

	leftArray := []rune{}
	rightArray := []rune{}

	for i, v := range s {
		if i < n {
			rightArray = append(rightArray, v)
		} else {
			leftArray = append(leftArray, v)
		}
	}
	return string(append(leftArray, rightArray...))
}
