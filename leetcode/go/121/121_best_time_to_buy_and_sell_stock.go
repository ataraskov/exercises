package main

// Problem: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

// test cases:
// [1,2,3] => 2
// [3,2,1] => 0
// [2,1,3] => 2
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	profit := 0
	min := prices[0]
	for _, price := range prices {
		if price < min {
			min = price
		} else {
			if profit < price-min {
				profit = price - min
			}
		}
	}
	return profit
}
