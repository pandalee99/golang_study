package main

func numWays(n int) int {
	ans := 0
	if n == 0 {
		return 1
	}
	if n == 1 || n == 2 {
		return n
	}
	first, second := 1, 2
	for i := 2; i < n; i++ {
		ans = (first + second) % 1000000007
		first = second
		second = ans
	}

	return ans
}
