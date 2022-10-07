package main

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	step_1 := 1
	step_2 := 2
	ans := 0
	for i := 2; i < n; i++ {
		ans = step_1 + step_2
		step_1 = step_2
		step_2 = ans
	}
	return ans
}
