package main

func minArray(numbers []int) int {
	temp, first, end := 0, 0, len(numbers)-1
	ans := numbers[end]
	//3, 3, 1, 3

	//1,3,3
	for first < end {
		//temp就是中间的值的坐标
		temp = (first + end) / 2
		//小小的判断
		if numbers[temp] < ans {
			ans = numbers[temp]
		}
		//看看中间的值会不会比end大，如果比end大则是旋转数组的尾部，
		if numbers[temp] > numbers[end] {
			first = temp + 1
			if first > end {
				break
			}
			//小小的判断
			if numbers[first] < ans {
				ans = numbers[first]
			}
		} else if numbers[temp] == numbers[end] {
			/*
				当等于的情况发生时，很难知道这到底是在数组的前面，还是在数组的后面，
				于是乎,要发射一个探测器，去判断这是数组的前面还是数组的后面
			*/
			detector := temp
			for numbers[detector] == numbers[end] && detector != end {
				detector++
			}
			//说明temp结点是尾部而不是头部，则执行尾部的逻辑
			if detector != end {
				first = temp + 1
				if first > end {
					break
				}
				//小小的判断
				if numbers[first] < ans {
					ans = numbers[first]
				}

			} else {
				//说明temp结点是头部而不是尾部，则执行头部的逻辑，注意的是，等于也算头部
				//小于则是旋转数组的头部
				end = temp - 1
				if first > end {
					break
				}
				//小小的判断
				if numbers[end] < ans {
					ans = numbers[end]
				}
			}
		} else {
			//小于则是旋转数组的头部
			end = temp - 1
			if first > end {
				break
			}
			//小小的判断
			if numbers[end] < ans {
				ans = numbers[end]
			}
		}

	}

	return ans
}
