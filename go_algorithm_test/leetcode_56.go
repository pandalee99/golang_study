package main

import "sort"

func leetcode_56_merge(intervals [][]int) [][]int {
	var ans [][]int

	//对二维切片进行排序
	var less func(i, j int) bool
	less = func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] <= intervals[j][1]
		}
		return intervals[i][0] <= intervals[j][0]
	}
	sort.Slice(intervals, less)
	//对二维切片进行排序，此时已然有序

	//然后先进行首次处理
	n := len(intervals[0])   //宽度
	max := intervals[0][n-1] //获得宽度的尾数，但此题目其实默认只有两个元素
	var temp []int
	temp = make([]int, 2)
	temp[0] = intervals[0][0]
	for i := 1; i < len(intervals); i++ {
		if max >= intervals[i][0] {
			//至少说明他们存在公共区间
			if max < intervals[i][1] {
				max = intervals[i][1]
			}
		} else {
			//则说明前面收集的区间的最大值仍然小于此区间的最小值
			//先将前面的值放入最终的ans
			temp[1] = max
			ans = append(ans, temp)
			//接着产生新逻辑
			max = intervals[i][1]
			temp = make([]int, 2)
			temp[0] = intervals[i][0]
		}
	}
	//收尾处理
	temp[1] = max
	ans = append(ans, temp)

	return ans
}
