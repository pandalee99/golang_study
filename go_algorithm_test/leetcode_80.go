package main

func removeDuplicates(nums []int) int {
	step := 1
	count := 1
	target := nums[0]
	//0,0,1,1,1,1,2,3,3
	for i := 1; i < len(nums); i++ {
		if nums[i] == target && count == 1 {
			nums[step] = target
			count++
			step++
			continue
		}
		if nums[i] == target && count == 2 {
			continue
		}
		if nums[i] != target && count <= 2 {
			count = 1
			nums[step] = nums[i]
			step++
			target = nums[i]
			continue
		}
		if nums[i] != target && count == 2 {
			count = 1
			nums[step] = nums[i]
			step++
			target = nums[i]
		}
	}
	return step
}
