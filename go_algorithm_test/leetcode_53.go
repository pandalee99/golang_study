package main

func maxSubArray(nums []int) int {
	previous := -1
	ans := nums[0]
	for i := 0; i < len(nums); i++ {
		if previous < 0 {
			previous = nums[i]
			if previous > ans {
				ans = previous
			}
		} else {
			previous += nums[i]
			if previous > ans {
				ans = previous
			}
		}
	}
	return ans
}
