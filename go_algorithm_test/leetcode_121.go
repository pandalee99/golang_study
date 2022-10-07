package main

func Leetcode_112_maxProfit(prices []int) int {
	ans := 0
	min := prices[0]

	for i := 1; i < len(prices); i++ {
		val := prices[i] - min
		if val > ans {
			ans = val
		}

		if prices[i] < min {
			min = prices[i]
		}
	}
	return ans
}
