package main

func maxArea(height []int) int {

	ans, head, tail := 0, 0, len(height)-1

	if height[0] > height[tail] {
		ans = height[tail] * (tail - head)
	} else {
		ans = height[head] * (tail - head)
	}
	for head != tail {

	}
	for head != tail {
		//补丁一
		if head+1 == tail-1 {
			if height[head+1] > ans {
				ans = height[head+1]
			}
			break
		}

		//前进
		if height[head+1] > height[tail-1] {
			head++
			if height[head] > height[tail] {
				if height[tail]*(tail-head) > ans {
					ans = height[tail] * (tail - head)
				}
			} else {
				if height[head]*(tail-head) > ans {
					ans = height[tail] * (tail - head)
				}
			}
		} else {
			tail--
			if height[head] > height[tail] {
				if height[tail]*(tail-head) > ans {
					ans = height[tail]*tail - head
				}
			} else {
				if height[head]*(tail-head) > ans {
					ans = height[tail]*tail - head
				}
			}
		}

	}

	return ans
}
