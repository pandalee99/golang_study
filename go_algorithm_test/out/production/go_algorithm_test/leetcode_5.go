package main

func longestPalindrome(s string) string {
	//判断到只有一个，则直接返回
	n := len(s)
	//if n == 1 {
	//	return s
	//}

	var ans string

	begin := 0
	max := 0
	//二维切片初始化
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	//a := make([][]int, leetcode_17_map) // 二维切片，3行
	//for i := range a {
	//	a[i] = make([]int, n) // 每一行4列
	//}

	//for i := 0; i < n; i++ {
	//	dp[0][i] = 1
	//}

	for i := 0; i < n; i++ {
		begin = 0 //对比的开始地址
		for j := i; j < n; j++ {
			if s[begin] == s[j] {
				if i-2 >= 0 { //dp现在处于大于第二行的的情况
					if dp[i-2][j-1] > 0 {
						dp[i][j] = dp[i-2][j-1] + 2
					} else {
						dp[i][j] = 0
					}
				} else { //dp现在处于小于或等于第二行的的情况
					if begin == j { //第一行的情况
						dp[i][j] = 1
					} else { //第二行的情况
						dp[i][j] = 2
					}
				}
			} else {
				dp[i][j] = 0
			}
			//截取字符串
			if dp[i][j] > max {
				ans = s[begin : j+1]
				max = dp[i][j]
			}
			begin++
		}

	}

	return ans
}

/*
acsa
  0 1 2 3
0 1 1 1 1
1 x 0 0 0
2 x x 0 0
3 x x x

abcba
  0 1 2 3 4
0 1 1 1 1 1
1 x 0 0 0 0
2 x x 0 3 0
3 x x x 0 0
4 x x x x 5


babad
  0 1 2 3 4
0 1 1 1 1 1
1 x 0 0 0 0
2 x x 3 3 0
3 x x x 0 0
4 x x x x 0

abbc
  0 1 2 3
0 1 1 1 1
1 x 0 2 0
2 x x 0 0
3 x x x 0

abba
  0 1 2 3
0 1 1 1 1
1 x 0 2 0
2 x x 0 0
3 x x x 2

*/
