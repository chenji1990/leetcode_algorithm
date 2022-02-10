package main

import (
	"fmt"
	"strconv"
)

func movingCount(m int, n int, k int) int {
	matrix := make([][]bool, m)
	for i := range matrix {
		matrix[i] = make([]bool, n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || (addBit(i)+addBit(j) > k) || matrix[i][j] {
			return 0
		}
		matrix[i][j] = true
		return dfs(i, j+1) + dfs(i, j-1) + dfs(i+1, j) + dfs(i-1, j) + 1
	}
	return dfs(0, 0)
}

func addBit(v int) int {
	src := strconv.Itoa(v)
	var sum int
	for i := 0; i < len(src); i++ {
		sum += (int(src[i]) - int('0'))
	}
	return sum
}

func main() {
	fmt.Println("res = ", movingCount(1, 2, 1))
}
