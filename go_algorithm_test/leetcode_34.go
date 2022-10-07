package main

func searchRange(nums []int, target int) []int {
	//把ans作为一个探测器，leetcode_39_ans[0]为left leetcode_39_ans[1]为right
	ans := []int{-1, -1}
	n := len(nums) - 1
	left, right := 0, n
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			//这里必然是等于target的，但是不知道边界，此时需要一个探测器去扫描
			ans[0] = mid
			for nums[ans[0]] == target && ans[0] > 0 {
				ans[0]--
			} //跳出循环后可能是：不等于target，或者ans[0]==0
			if nums[ans[0]] != target {
				ans[0]++
			}
			//处理右边
			ans[1] = mid
			for nums[ans[1]] == target && ans[1] < n {
				ans[1]++
			} //跳出循环后可能是：不等于target，或者ans[1]==leetcode_46_nums_length-1
			if nums[ans[1]] != target {
				ans[1]--
			}
			break

		}
	}

	return ans
}
