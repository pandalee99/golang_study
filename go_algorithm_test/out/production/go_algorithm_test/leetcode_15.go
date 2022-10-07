package main

import "sort"

func threeSum(nums []int) [][]int {
	var ans [][]int
	if len(nums) < 3 {
		return nil
	}
	target, first, end := 0, 0, 0
	//排序
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		//若此数等于它的前一个数，则跳过？以此来达到去重的目的
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target = i
		first = i + 1
		end = len(nums) - 1
		for first < end {
			if nums[target]+nums[first]+nums[end] == 0 {
				//这是一个在二维数组中塞一维数组的方法
				arr := make([]int, 3)
				arr[0] = nums[target]
				arr[1] = nums[first]
				arr[2] = nums[end]
				ans = append(ans, arr)
				//去重
				for first < end && nums[first] == nums[first+1] {
					first++
				}
				for first < end && nums[end] == nums[end-1] {
					end--
				}
				first++
				end--
				continue
			}
			if nums[target]+nums[first]+nums[end] > 0 {
				end--
				continue
			}
			if nums[target]+nums[first]+nums[end] < 0 {
				first++
				continue
			}
		}

	}
	return ans
}

/*
	leetcode_17_map := 2
	n := 2
	//var dp [leetcode_17_map][n]int  这样会报错提示无法使用变量
	var dp [][]int
	for x := 0; x < leetcode_17_map+1; x++ {  //循环为一维长度
	arr := make([]int, n+1) //创建一个一维切片
	dp = append(dp, arr)    //把一维切片,当作一个整体传入二维切片中
	}
*/
