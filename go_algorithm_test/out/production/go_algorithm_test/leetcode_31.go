package main

func nextPermutation(nums []int) {
	n := len(nums)
	if n == 1 {
		return
	}
	first, second := n-1, n-1
	for first > 0 {
		if nums[first] > nums[first-1] {
			first = first - 1
			break
		}
		first--
	}
	for second > 0 {
		if nums[second] > nums[first] {
			//交换
			temp := nums[first]
			nums[first] = nums[second]
			nums[second] = temp
			break
		}
		second--
	}
	if first != 0 {
		first++
	}

	second = n - 1
	for first < second {
		//交换
		temp := nums[first]
		nums[first] = nums[second]
		nums[second] = temp
		first++
		second--
	}
	/*
	   	//end表示数组最后一个元素
	   	//[5,4,6,3,2,1]
	   	//[5,6,1,2,3,4]

	   	//[1,3,2]
	   	//[2,1,3]

	   	//
	   	end := n - 1
	   	exchangePointer := end
	   	for exchangePointer > 0 {
	   		for i := exchangePointer - 1; i >= 0; i-- {
	   			if nums[exchangePointer] > nums[i] {
	   				//交换
	   				temp := nums[exchangePointer]
	   				nums[exchangePointer] = nums[i]
	   				nums[i] = temp
	   				exchangePointer = i + 1
	   				goto part2
	   			}
	   		}
	   		exchangePointer--
	   	}
	   part2:
	   	for exchangePointer < end {
	   		temp := nums[exchangePointer]
	   		nums[exchangePointer] = nums[end]
	   		nums[end] = temp
	   		exchangePointer++
	   		end--
	   	}
	   	for i := range nums {
	   		fmt.Print(nums[i])
	   	}

	*/
}
