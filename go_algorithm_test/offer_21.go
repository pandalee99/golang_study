package main

func exchange(nums []int) []int {
	odd := 0
	n := len(nums)
	for i := 0; i < n; i++ {

		if nums[i]&1 == 0 {
			//如果是偶数
		} else {
			//如果是奇数
			//交换
			nums[i], nums[odd] = nums[odd], nums[i]
			//奇数指数增加
			odd++
		}

	}
	return nums
}
