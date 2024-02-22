package main

// Problem: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

// test cases:
// [1,2,3] => 2
// [3,2,1] => 0
// [2,1,3] => 2
func maxProfit(prices []int) int {
	max := func(prices []int) int {
		max := 0
		for _, n := range prices {
			if n > max {
				max = n
			}
		}
		return max
	}

	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		n := max(prices[i+1:]) - prices[i]
		if n > profit {
			profit = n
		}
	}
	return profit
}
