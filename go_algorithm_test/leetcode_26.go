package main

func Leetcode_26_removeDuplicates(nums []int) int {
	//考虑特殊的情况
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	ans := 1
	//使用双指针法即可
	first := 0
	//i为second指针
	for i := 1; i < n; i++ {
		if nums[i] == nums[first] {
			continue
		} else {
			first++
			ans++
			nums[first] = nums[i]
		}

	}

	return ans
}
