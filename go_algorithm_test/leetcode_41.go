package main

func firstMissingPositive(nums []int) int {
	n := len(nums)
	//先让数组归位，
	//如果小于等于0则不管它
	//如果大于0但是也大于n，也不管
	for i := 0; i < n; {
		if nums[i] > 0 && nums[i] <= n {
			//交换位置
			temp := nums[i]
			nums[i] = nums[nums[i]-1]
			nums[temp-1] = temp
			//看看交换后的值是否需要再正确的归位
			if nums[i] > 0 && nums[i] <= n {
				//如果归位的值已经被正确归位了，则跳过
				if nums[i] == nums[nums[i]-1] {
					i++
				}
				//归位后的值没有正确被归位，需要再次归位循环
				continue
			} else {
				//非需归位值，跳过
				i++
			}
		} else {
			//非需归位值，跳过
			i++
		}

	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}
