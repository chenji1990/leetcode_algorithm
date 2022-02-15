package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	row, column := len(matrix), len(matrix[0])
	if row == 1 {
		return matrix[0]
	}
	if column == 1 {
		res := []int{}
		for j := 0; j < row; j++ {
			res = append(res, matrix[j][0])
		}
		return res
	}

	res := []int{}
	subMatrix := [][]int{}

	i, j := 0, 0
	for j = 0; j < column; j++ {
		res = append(res, matrix[i][j])
	}

	j = column - 1
	for i = 1; i < row-1; i++ {
		res = append(res, matrix[i][j])
		subMatrix = append(subMatrix, matrix[i][1:j])
	}

	i = row - 1
	for j = column - 1; j > -1; j-- {
		res = append(res, matrix[i][j])
	}

	j = 0
	for i = row - 2; i > 0; i-- {
		res = append(res, matrix[i][j])
	}

	return append(res, spiralOrder(subMatrix)...)
}

// 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
// 示例 1：
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]

// 示例 2：
// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
