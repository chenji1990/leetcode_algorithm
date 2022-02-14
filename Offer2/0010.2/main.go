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
