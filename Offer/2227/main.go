package main

func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	const mod int = 1e9 + 7
	v1, v2 := 1, 2
	var i int
	for i = 2; i < n; i++ {
		v1, v2 = v2, (v1+v2)%mod
	}
	return v2
}

// 一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。

// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
