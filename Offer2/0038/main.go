package main

import "sort"

func permutation(s string) (res []string) {
	array := []byte(s)
	sort.Slice(array, func(i, j int) bool { return array[i] < array[j] })
	length := len(array)
	temp := make([]byte, 0, length)
	vis := make([]bool, length)
	var backtrack func(int)
	backtrack = func(i int) {
		if i == length {
			res = append(res, string(temp))
			return
		}
		for j, yes := range vis {
			if yes || (j > 0 && !vis[j-1] && array[j-1] == array[j]) {
				continue
			}
			vis[j] = true
			temp = append(temp, array[j])

			backtrack(i + 1)

			vis[j] = false
			temp = temp[:len(temp)-1]
		}
	}
	backtrack(0)
	return
}
