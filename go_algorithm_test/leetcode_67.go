package main

func addBinary(a string, b string) string {
	ans := ""
	i := len(a) - 1
	j := len(b) - 1
	flag := false
	var bytes []byte
	for i >= 0 && j >= 0 {
		if flag == false {
			if a[i] == '1' && b[j] == '1' {
				bytes = append([]byte{'0'}, bytes...)
				flag = true
			} else if a[i] == '0' && b[j] == '1' {
				bytes = append([]byte{'1'}, bytes...)
			} else if a[i] == '1' && b[j] == '0' {
				bytes = append([]byte{'1'}, bytes...)
			} else if a[i] == '0' && b[j] == '0' {
				bytes = append([]byte{'0'}, bytes...)
			}
		} else {
			if a[i] == '1' && b[j] == '1' {
				bytes = append([]byte{'1'}, bytes...)
			} else if a[i] == '0' && b[j] == '1' {
				bytes = append([]byte{'0'}, bytes...)
			} else if a[i] == '1' && b[j] == '0' {
				bytes = append([]byte{'0'}, bytes...)
			} else if a[i] == '0' && b[j] == '0' {
				flag = false
				bytes = append([]byte{'1'}, bytes...)
			}
		}
		i--
		j--
	}
	for i == -1 && j != -1 {
		if flag == false {
			if b[j] == '1' {
				bytes = append([]byte{'1'}, bytes...)
			} else if b[j] == '0' {
				bytes = append([]byte{'0'}, bytes...)
			}
		} else {
			if b[j] == '1' {
				bytes = append([]byte{'0'}, bytes...)
			} else if b[j] == '0' {
				bytes = append([]byte{'1'}, bytes...)
				flag = false
			}
		}
		j--
	}
	if i == -1 && j == -1 && flag {
		bytes = append([]byte{'1'}, bytes...)
		ans = string(bytes)

		return ans

	}
	for j == -1 && i != -1 {
		if flag == false {
			if a[i] == '1' {
				bytes = append([]byte{'1'}, bytes...)
			} else if a[i] == '0' {
				bytes = append([]byte{'0'}, bytes...)
			}
		} else {
			if a[i] == '1' {
				bytes = append([]byte{'0'}, bytes...)
			} else if a[i] == '0' {
				bytes = append([]byte{'1'}, bytes...)
				flag = false
			}
		}
		i--
	}
	if i == -1 && j == -1 && flag {
		bytes = append([]byte{'1'}, bytes...)
	}

	ans = string(bytes)

	return ans
}
