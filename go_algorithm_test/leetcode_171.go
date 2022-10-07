package main

func titleToNumber(columnTitle string) int {
	n := len(columnTitle) - 1
	ans := 0
	for i := 0; i < len(columnTitle); i++ {
		var temp int
		temp = int(columnTitle[i] - 64)
		for j := 0; j < n; j++ {
			temp *= 26
		}
		n--
		ans += temp
	}

	return ans
}
