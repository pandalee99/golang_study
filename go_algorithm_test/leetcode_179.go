package main

/*
未解之谜


import (
	"sort"
	"strings"
)

func largestNumber(nums []int) string {
	var less func(i, j int) bool
	less = func(i, j int) bool {
		a, b := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		n, m := len(a), len(b)
		y := 0
		for x := 0; x < n; {
			if y == m {
				for i := 0; i < m; i++ {
					if b[i] > a[n-1] {
						return false
					} else if b[i] == a[n-1] {
						continue
					} else {
						//else if a[i]>b[n]
						return true
					}
				}

				//if a[n-1] < b[n] {
				//	return false
				//} else {
				//	return true
				//}
				return false
			}
			if a[x] == b[y] {
				x++
				y++
			} else if a[x] < b[y] {
				return false
			} else {
				return true
			}
		}
		if m == n {
			return false
		} else {
			//m>n 3,34
			for i := 0; i < n; i++ {
				if a[i] < b[m-1] {
					return false
				} else if a[i] == b[m-1] {
					continue
				} else {
					//else if a[i]>b[n]
					return true
				}
			}

			//if a[n-1] < b[n] {
			//	return false
			//} else {
			//	return true
			//}
			return false
		}

	}
	sort.Slice(nums, less)
	var build strings.Builder
	if nums[0] == 0 {
		return "0"
	}
	for i := 0; i < len(nums); i++ {
		build.WriteString(strconv.Itoa(nums[i]))
	}
	return build.String()
}
*/
//3432 3,3432
//343234323
//343233432
//34 3
// 44934 4493
//4493 44934
//44934 4493

/*
https://leetcode.cn/problems/largest-number/
https://zhuanlan.zhihu.com/p/352301583

给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。


示例 1：

输入：nums = [10,2]
输出："210"


输入：nums = [3,30,34,5,9]
输出："9534330"
*/
