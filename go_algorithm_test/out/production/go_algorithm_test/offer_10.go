package main

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	ans := 0
	one, two := 0, 1
	for i := 1; i < n; i++ {
		ans = (one + two) % 1000000007
		one = two
		two = ans
	}

	return ans
}
