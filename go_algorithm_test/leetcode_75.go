package main

func sortColors(nums []int) {
	first := 0
	mid := 0
	end := len(nums) - 1
	for mid <= end {
		for nums[mid] != 1 && mid >= first && mid <= end {
			if nums[mid] == 0 {
				nums[first], nums[mid] = nums[mid], nums[first]
				first++

			}
			if nums[mid] == 2 {
				nums[end], nums[mid] = nums[mid], nums[end]
				end--
			}
		}
		mid++
	}
}
