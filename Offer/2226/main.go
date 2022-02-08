package main

func fib(n int) int {
	if n <= 1 {
		return n
	}
	const mod int = 1e9 + 7
	var i int
	var v1, v2 = 0, 1
	for i = 1; i < n; i++ {
		v1, v2 = v2, (v1+v2)%mod
	}
	return v2
}
