package main

func leetcode_122_maxProfit(prices []int) int {
	temp := 0
	min := prices[0]
	max := 0

	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < max {
			ans += temp
			temp = 0
			min = prices[i]
			max = 0
			continue
		}

		val := prices[i] - min
		if val > temp {
			temp = val
			max = prices[i]
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	ans += temp

	return ans
}
