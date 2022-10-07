package main

func convertToTitle(columnNumber int) string {
	var ans string
	for columnNumber > 26 {
		if columnNumber%26 == 0 {
			ans = "Z" + ans
			columnNumber /= 26
			columnNumber--
		} else {
			//从头追加
			ans = string(columnNumber%26+64) + ans
			columnNumber /= 26
		}
	}
	ans = string(columnNumber+64) + ans

	return ans
}
