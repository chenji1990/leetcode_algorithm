package main

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	array := []map[int]bool{
		{},
		{},
	}

	i, j, index := 0, 0, 0
	row, column := len(matrix), len(matrix[0])

	for i = 0; i < row; i++ {
		for j = 0; j < column; j++ {
			if matrix[i][j] == 0 {
				array[0][i], array[1][j] = true, true
			}
		}
	}

	for i = range array[0] {
		for index = 0; index < column; index++ {
			matrix[i][index] = 0
		}
	}

	for j = range array[1] {
		for index = 0; index < row; index++ {
			matrix[index][j] = 0
		}
	}
}

// fmt.Println(matrix)

func main() {
	array := [][]int{
		// {1, 1, 1},
		// {1, 0, 1},
		// {1, 1, 1},
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(array)
}
