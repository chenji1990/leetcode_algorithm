package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 || r*c != len(mat)*len(mat[0]) {
		return mat
	}

	res := [][]int{}
	newRowArray := []int{}
	column := 0

	for _, rowArray := range mat {
		for _, value := range rowArray {
			if column < c {
				newRowArray = append(newRowArray, value)
				column += 1
			} else {
				res = append(res, newRowArray)
				newRowArray = []int{value}
				column = 1
			}
		}
	}
	return append(res, newRowArray)
}

func main() {
	matrixReshape([][]int{
		{1, 2}, {3, 4},
	}, 1, 4)
}
