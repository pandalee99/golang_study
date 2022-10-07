package main

func partition(s string) [][]string {
	var ans [][]string

	n := len(s)
	//二维切片初始化
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	//辅助函数
	isPalindrome := func(first, end int) bool {
		if first >= end {
			return true
		} else {
			return dp[first][end]
		}

	}
	var partitionHelper func(temp []string, s string, begin int)
	partitionHelper = func(temp []string, s string, begin int) {
		if begin == n {
			final := make([]string, len(temp))
			copy(final, temp)
			ans = append(ans, final)
			return
		}
		for i := begin; i < n; i++ {
			if i == begin {
				dp[i][i] = true
				temp = append(temp, s[i:i+1])
				partitionHelper(temp, s, i+1)
				temp = temp[:len(temp)-1]
			} else {
				if s[begin] == s[i] && isPalindrome(begin+1, i-1) {
					dp[begin][i] = true
					temp = append(temp, s[begin:i+1])
					partitionHelper(temp, s, i+1)
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		var temp []string
		//主要处理逻辑
		if i == 0 {
			dp[0][i] = true
			temp = append(temp, s[:i+1])
			//分割函数
			partitionHelper(temp, s, i+1)
		} else {
			if s[0] == s[i] && isPalindrome(0+1, i-1) {
				//首先更新dp
				dp[0][i] = true
				//其次放入temp
				temp = append(temp, s[:i+1])
				//分割函数
				partitionHelper(temp, s, i+1)
			}
		}
	}

	return ans
}
