package main

func hammingWeight(num uint32) int {
	count := 0
	for i := 0; i < 32; i++ {
		if num == 0 {
			return count
		}
		if num%2 == 1 {
			count += 1
		}
		num >>= 1
	}
	return count
}
