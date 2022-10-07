package main

import "fmt"

func trailingZeroes(n int) int {
	count := 0
	sum := 1
	for i := 2; i <= n; i++ {
		sum *= i
		fmt.Println(sum)
		if sum%1000000 == 0 {
			count += 6
			sum /= 1000000
		} else if sum%100000 == 0 {
			count += 5
			sum /= 100000
		} else if sum%10000 == 0 {
			count += 4
			sum /= 10000
		} else if sum%1000 == 0 {
			count += 3
			sum /= 1000
		} else if sum%100 == 0 {
			count += 2
			sum /= 100
		} else if sum%10 == 0 {
			count++
			sum /= 10
		}
		if sum >= 100000 {
			sum = sum - (sum/100000)*100000
		}
		fmt.Println(sum)
	}
	return count
}
