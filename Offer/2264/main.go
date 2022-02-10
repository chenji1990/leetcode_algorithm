package main

import "fmt"

func maxValue(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	row, column := len(grid), len(grid[0])
	var i, j int
	for i = 1; i < row; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j = 1; j < column; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i = 1; i < row; i++ {
		for j = 1; j < column; j++ {
			grid[i][j] += max(grid[i-1][j], grid[i][j-1])
		}
	}
	fmt.Println(grid)
	return grid[row-1][column-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
