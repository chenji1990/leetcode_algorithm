package main

import "math"

func maxProfit(prices []int) int {
	var minPrice int = math.MaxInt64
	var price int = math.MaxInt64
	var profit int = 0
	var tempProfit int = 0
	for _, price = range prices {
		tempProfit = price - minPrice
		if tempProfit < 0 {
			minPrice = price
		} else if tempProfit > 0 && tempProfit > profit {
			profit = tempProfit
		}
	}
	return profit
}
