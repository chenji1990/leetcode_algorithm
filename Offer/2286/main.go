package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var minValue int = 1e9
	var profit, temp int = 0, 0

	for _, price := range prices {
		temp = price - minValue
		if temp > profit {
			profit = temp
		}
		if price < minValue {
			minValue = price
		}
	}
	return profit
}
