package main

// 动态规划
func cuttingRope(n int) int {
	array := make([]int, n+1)
	array[0], array[1], array[2] = 0, 1, 1
	j, half := 0, 0
	for i := 3; i <= n; i++ {
		array[i] = array[i-1]
		half = i / 2
		for j = 1; j <= half; j++ {
			array[i] = max(array[i], max(array[j], j)*max(array[i-j], i-j))
		}
	}
	return array[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。
// 请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。
