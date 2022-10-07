package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	//排除特殊情况
	if m == 0 {
		copy(nums1, nums2)
	}
	end := m + n - 1
	target := m - 1
	i := 0
	for i = n - 1; i >= 0 && target >= 0; {
		/*
					[0,1,2,3,7,9,0,0,0,0,0,0]
					6
					[2,4,5,6,8,11]
					6
					0,1,2,

				输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
				输出：[1,2,2,3,5,6]

			[3,3,3,0,0]
			3
			[2,4]
			2
			3334
		*/

		if nums2[i] >= nums1[target] {
			nums1[end] = nums2[i]
			i--
		} else {
			nums1[end] = nums1[target]
			target--
		}
		end--

	}
	//特殊情况
	if target == -1 && end != -1 {
		for end != -1 {
			nums1[end] = nums2[i]
			i--
			end--
		}
	}

}
