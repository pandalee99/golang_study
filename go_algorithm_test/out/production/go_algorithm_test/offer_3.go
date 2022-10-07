package main

func findRepeatNumber(nums []int) int {
	hashmap := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if hashmap[nums[i]] == 1 {
			return nums[i]
		}
		hashmap[nums[i]] = 1
	}

	return 0
}
