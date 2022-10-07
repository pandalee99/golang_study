package main

func jump(nums []int) int {
	ans := 0
	max := -1
	for i := 0; i < len(nums)-1; {
		max = -1
		for j := 1; j <= nums[i]; j++ {
			if i+j == len(nums)-1 {
				return ans + 1
				break
			}
			if max == -1 || nums[i+j]+i+j-max >= nums[max] {
				max = i + j
			}
		}
		i = max
		ans++
	}

	return ans
}

/*
	stepArr := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		stepArr[i] = -1
	}
	stepArr[0] = 0
	//step := 0
	for i := 0; i < len(nums); i++ {
		//对数组的每个值进行探查
		for j := 1; j <= nums[i]; j++ {
			//边界条件
			if i+j == len(nums) {
				break
			}
			if stepArr[i]+1 < stepArr[i+j] || stepArr[i+j] == -1 {
				stepArr[i+j] = stepArr[i] + 1
			}
		}
	}
	return stepArr[len(nums)-1]

*/
