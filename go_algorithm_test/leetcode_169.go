package main

func majorityElement(nums []int) int {
	ans, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if ans == nums[i] {
			count++
		} else {
			count--
			if count == 0 {
				i++
				count++
				ans = nums[i]
			}
		}
	}
	return ans
}
