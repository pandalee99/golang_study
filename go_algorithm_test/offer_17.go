package main

func printNumbers(n int) []int {
	var ans []int
	size := 10
	for i := 1; i < n; i++ {
		size = size * 10
	}
	for i := 1; i < size; i++ {
		ans = append(ans, i)
	}

	return ans
}
