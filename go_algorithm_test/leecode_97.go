package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	//isInterleave("aabcc", "dbbca", "aadbbbcacc")
	m, n := len(s1), len(s2)
	if (m + n) != len(s3) {
		return false
	}
	if s1 == "" {
		return s2 == s3
	}
	if s2 == "" {
		return s1 == s3
	}
	judge := make([][]bool, m+1)
	for i := range judge {
		judge[i] = make([]bool, n+1)
	}
	//初始化map
	for i := 0; i < m; i++ {
		if s1[i] == s3[i] {
			judge[i+1][0] = true
		} else {
			break
		}
	}
	for i := 0; i < n; i++ {
		if s2[i] == s3[i] {
			judge[0][i+1] = true
		} else {
			break
		}
	}
	//开始逻辑
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			judge[i][j] = (judge[i-1][j] && s1[i-1] == s3[i+j-1]) || (judge[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return judge[m][n]
}

/*
虽败犹荣
最后一个用例没有通过
已通过：105/106



func isInterleave(s1 string, s2 string, s3 string) bool {

	flag = false
	m, n, length := len(s1), len(s2), len(s3)
	if (m + n) != length {
		return false
	}
	p1, p2 := 0, 0
	for i := 0; i < length; i++ {
		if p1 == m {
			for p2 != n {
				if s2[p2] != s3[i] {
					return false
				}
				p2++
				i++
			}
			return true
		}
		if p2 == n {
			for p1 != m {
				if s1[p1] != s3[i] {
					return false
				}
				p1++
				i++
			}
			return true
		}

		if s1[p1] == s3[i] && s2[p2] != s3[i] {
			p1++
			continue
		}
		if s1[p1] != s3[i] && s2[p2] == s3[i] {
			p2++
			continue
		}
		if s1[p1] == s3[i] && s2[p2] == s3[i] {
			//
			isInterleaveFunction(s1[p1+1:], s2[p2:], s3[p1+p2+1:])
			isInterleaveFunction(s1[p1:], s2[p2+1:], s3[p1+p2+1:])
			return flag
		}
		if s1[p1] != s3[i] && s2[p2] != s3[i] {
			return false
		}
	}
	flag = true
	return flag
}

func isInterleaveFunction(s1 string, s2 string, s3 string) {

	if flag == true {
		return
	}
	m, n, length := len(s1), len(s2), len(s3)
	p1, p2 := 0, 0
	for i := 0; i < length; i++ {
		if p1 == m {
			for p2 != n {
				if s2[p2] != s3[i] {

					return
				}
				p2++
				i++
			}

			flag = true
			return
		}
		if p2 == n {
			for p1 != m {
				if s1[p1] != s3[i] {

					return
				}
				p1++
				i++
			}
			flag = true
			return
		}

		if s1[p1] == s3[i] && s2[p2] != s3[i] {
			p1++
			continue
		}
		if s1[p1] != s3[i] && s2[p2] == s3[i] {
			p2++
			continue
		}
		if s1[p1] == s3[i] && s2[p2] == s3[i] {
			//
			isInterleaveFunction(s1[p1+1:], s2[p2:], s3[p1+p2+1:])
			isInterleaveFunction(s1[p1:], s2[p2+1:], s3[p1+p2+1:])
			return
		}
		if s1[p1] != s3[i] && s2[p2] != s3[i] {
			return
		}
	}

	flag = true

	return

}

var flag bool
*/
