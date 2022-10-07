package main

func removeElement(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	//双指针法
	ans := 0
	first := 0
	for i := 0; i < n; i++ {
		if nums[i] == val {
			continue
		} else {
			ans++
			nums[first] = nums[i]
			first++
		}
	}

	return ans
}
