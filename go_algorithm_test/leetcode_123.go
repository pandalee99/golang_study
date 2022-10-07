package main

func maxProfit(prices []int) int {
	temp := 0
	//min := prices[0]
	max := 0

	first := 0
	second := 0
	//[2,7,1,10,4,3,9]
	for i := 1; i < len(prices); i++ {
		val := prices[i] - prices[i-1]
		temp = temp + val
		
		if temp > 0 {
			//预估收益还是正的，亏的起
			if temp > max {
				max = temp
			}
		} else {
			//再继续放长线钓大鱼已经没有意义了，及时收尾
			//找出其中出现过的最大的值，进行对比

			//比最大的还要大，那么顺位继承
			if max > first {
				second = first
				first = max
			} else if max > second {
				//只比第二大的大，那么就换掉第二大的值
				second = max
			}
			//重新做人
			temp = 0
			max = 0
		}

	}
	//补充
	//比最大的还要大，那么顺位继承
	if max > first {
		second = first
		first = max
	} else if max > second {
		//只比第二大的大，那么就换掉第二大的值
		second = max
	}

	return first + second
}
