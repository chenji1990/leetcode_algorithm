package main

func isValidSudoku(board [][]byte) bool {
	var array = [][][]int{}
	var index byte
	for index = 1; index <= 9; index++ {
		array = append(array, [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		})
	}

	var i, j, section int

	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			index = board[i][j] - '1'
			if array[index][0][i] > 0 || array[index][1][j] > 0 {
				return false
			}

			section = (i/3)*3 + j/3
			if array[index][2][section] > 0 {
				return false
			}
			array[index][0][i] = 1
			array[index][1][j] = 1
			array[index][2][section] = 1
		}
	}
	return true
}
