package main

func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	//先找到最小值点，然后再二分
	temp, first, end := 0, 0, len(nums)-1
	//初始最小值的坐标
	min := 0
	for first < end {
		//temp就是中间的值的坐标
		temp = (first + end) / 2
		//判断
		if nums[temp] < nums[min] {
			min = temp
		}
		if nums[temp] > nums[end] {
			first = temp + 1
			if first > end {
				break
			}
			//判断
			if nums[first] < nums[min] {
				min = first
			}
		} else {
			end = temp - 1
			if first > end {
				break
			}
			//判断
			if nums[end] < nums[min] {
				min = end
			}
		}
	}
	//这样就找到了起点的坐标
	//然后正常二分查找
	//先确立好坐标
	if target <= nums[len(nums)-1] {
		first = min
		end = len(nums) - 1
	} else {
		first = 0
		if min != 0 {
			end = min - 1
		} else {
			end = 0
		}
	}
	if nums[first] == target {
		return first
	}
	for first < end {
		temp = (first + end) / 2
		if nums[temp] == target {
			return temp
		}
		if target > nums[temp] {
			first = temp + 1
			if first > end {
				break
			}
			if nums[first] == target {
				return first
			}
		} else {
			end = end - 1
			if first > end {
				break
			}
			if nums[end] == target {
				return end
			}
		}
	}
	return -1
}
