package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n > 0 {
		val := x
		for i := 1; i < n; i++ {
			x = x * val
		}
		return x
	} else {
		val := x
		for i := 0; i < -n; i++ {
			x = x * val
		}
		return 1 / x
	}

}
