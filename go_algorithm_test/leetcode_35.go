package main

func searchInsert(nums []int, target int) int {
	//表示坐标
	ans := 0
	left, right := 0, len(nums)-1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			//leetcode_39_ans = right
			right = mid - 1
		} else {
			left = mid + 1
			ans = left
		}
	}

	return ans
}
