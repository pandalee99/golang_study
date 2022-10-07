package main

/*
		这题与其说是 搜索题，不如说是找规律题，

		例如
	[1,2,3,4,5]
	[3,4,5,1,2]

	可以产生一个代价数组 res=[-2,-2,-2,3,3]
	之后将数组值累加
	最后只需要找到 累加值的最低点 的后一个点就可以
*/
func canCompleteCircuit(gas []int, cost []int) int {

	n := len(gas)
	min := gas[0] - cost[0]
	sum := gas[0] - cost[0]
	index := 0
	for i := 1; i < n; i++ {
		sum += gas[i] - cost[i]
		if sum < min {
			min = sum
			index = i
		}
	}

	if sum >= 0 {
		return (index + 1) % n
	} else {
		return -1
	}
}

/*


func canCompleteCircuit(gas []int, cost []int) int {

	n := len(gas)
	if n == 1 {
		if gas[0]-cost[0] >= 0 {
			return 0
		} else {
			return -1
		}
	}

	for i := 0; i < n; i++ {
		if gas[i]-cost[i] < 0 {
			continue
		}
		temp := gas[i]
		j := i
		count := n
		for ; count != 0; count-- {

			temp -= cost[j]
			if temp < 0 {
				break
			}
			j++
			if j == n {
				j = 0
			}
			temp += gas[j]

		}
		if temp >= 0 {
			return i
		}
	}

	return -1
}
*/
