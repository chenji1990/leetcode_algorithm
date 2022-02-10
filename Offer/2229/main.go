package main

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(board)*len(board[0]) < len(word) {
		return false
	}
	row, column, length := len(board), len(board[0]), len(word)
	var ok bool
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if k >= length {
			return true
		}
		if i >= 0 && i < row && j >= 0 && j < column && board[i][j] == word[k] {
			board[i][j] = '0'
			ok = dfs(i-1, j, k+1) || dfs(i+1, j, k+1) || dfs(i, j-1, k+1) || dfs(i, j+1, k+1)
			board[i][j] = word[k]
			return ok
		}
		return false
	}
	var i, j int
	for i = 0; i < row; i++ {
		for j = 0; j < column; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func exist2(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	if n*m < len(word) {
		return false
	}
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		// 终止条件
		if k >= len(word) {
			return true
		}
		// 终止条件，坐标范围错误或者 i j 坐标对应 board 的值不等于 k 坐标对应 word 的值
		if i < 0 || j < 0 || i >= n || j >= m || board[i][j] != word[k] {
			return false
		}
		// 如果往回前找，不会找到相同字符，如 ABAB，k = 2 时，往前往后都是 B ；
		// 将 B 修改（剪枝）为不存在的字符，杜绝往回找成功的可能性。
		board[i][j] = '0'
		res := dfs(i, j-1, k+1) || dfs(i, j+1, k+1) || dfs(i+1, j, k+1) || dfs(i-1, j, k+1)
		// 找完了就改回来
		board[i][j] = word[k]
		return res
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
