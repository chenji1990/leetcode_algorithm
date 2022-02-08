package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row, column := len(matrix), len(matrix[0])
	var i, j, value int
	for i = 0; i < row; i++ {
		for j = 0; j < column; j++ {
			value = matrix[i][j] - target
			if value == 0 {
				return true
			} else if value < 0 {
				continue
			} else {
				break
			}
		}
	}
	return false
}
