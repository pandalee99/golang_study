package main

func canJump(nums []int) bool {
	max := -1
	for i := 0; i < len(nums)-1; {
		max = -1
		//如果选中的最大值等于0，那么肯定是到不了的
		if nums[i] == 0 {
			return false
		}
		for j := 1; j <= nums[i]; j++ {
			//说明能够到达尾部
			if i+j == len(nums)-1 {
				return true
			}
			if max == -1 || nums[i+j]+i+j-max >= nums[max] {
				max = i + j
			}
		}
		i = max
	}

	return true
}
