package main

func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	n--
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}
	if nums[n] > nums[n-1] {
		return n
	}
	return 0
}
