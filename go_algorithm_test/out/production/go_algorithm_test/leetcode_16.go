package main

import "sort"

func threeSumClosest(nums []int, target int) int {

	mid, first, end := 0, 0, 0
	//排序
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		mid = i
		first = i + 1
		end = len(nums) - 1
		for first < end { //-3 0 1 2
			if threeSumClosest_abs(nums[mid]+nums[first]+nums[end]-target) <= threeSumClosest_abs(ans-target) {
				ans = nums[mid] + nums[first] + nums[end]
				continue
			}
			if nums[mid]+nums[first]+nums[end] > target {
				end--
				continue
			}
			if nums[mid]+nums[first]+nums[end] < target {
				first++
				continue
			}
			return ans
		}

	}
	return ans
}

func threeSumClosest_abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}
