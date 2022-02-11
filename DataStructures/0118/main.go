package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int{
			{1},
		}
	}
	if numRows == 2 {
		return [][]int{
			{1},
			{1, 1},
		}
	}

	res := [][]int{
		{1},
		{1, 1},
	}
	rowArray := []int{1}
	var row, column int
	for row = 2; row < numRows; row++ {
		for column = 1; column < row; column++ {
			rowArray = append(rowArray, res[row-1][column-1]+res[row-1][column])
		}
		res = append(res, append(rowArray, 1))
		rowArray = []int{1}
	}
	return res
}

func main() {
	fmt.Println(generate(10))
}
