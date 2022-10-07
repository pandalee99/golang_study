package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	var a strings.Builder
	mutex1 := false
	mutex2 := false
	var b strings.Builder
	n, m := len(version1), len(version2)
	i, j := 0, 0
	for ; i < n; i++ {
		if version1[i] == '.' {
			//如果找到了结束点，那么可以去和别的有效值去进行比较了
			for ; j < m; j++ {
				if version2[j] == '.' {
					j++
					mutex2 = false
					break
				} else if version2[j] == '0' {
					if mutex2 == false {
						continue
					} else {
						b.WriteByte(version2[j])
					}
				} else {
					mutex2 = true
					b.WriteByte(version2[j])
				}

			}
			temp1, _ := strconv.Atoi(a.String())
			temp2, _ := strconv.Atoi(b.String())
			a.Reset()
			b.Reset()
			if temp1 > temp2 {
				return 1
			} else if temp1 < temp2 {
				return -1
			}
			//如果都没有返回，那就相等了，继续判断
			mutex1 = false
		} else if version1[i] == '0' {
			if mutex1 == false {
				continue
			} else {
				a.WriteByte(version1[i])
			}
		} else {
			mutex1 = true
			a.WriteByte(version1[i])
		}

	}
	//到了这一步，v1已经进行完了
	if mutex1 == false {
		//如果锁没有打开，说明v1根本就没有值，这时候只需要看V2有没有值就行了
		//如果V2也到了，那么就没值了，两者相等
		if j == m {
			return 0
		}
		//继续将V2进行到底
		for ; j < m; j++ {
			//换一个角度想，如果V2即不等于'.'又不等于'0',那么是肯定大于v1的
			if version2[j] == '.' || version2[j] == '0' {
				continue
			} else {
				return -1
			}
			/*


				if version2[j]=='.'{
					//这时候和上面的有所不同
					if b.String()==""{
						continue
					}else {
						return -1
					}
				}else if version2[j]=='0'{
					if mutex2==false{
						continue
					}else {
						b.WriteByte(version2[j])
					}
				}else {
					mutex2=true
					b.WriteByte(version2[j])
				}

			*/
		}
		//如果没有返回值，那么只能是相等了
		return 0
		/*
			if mutex2==false{
				//还是没有开锁，说明仍然为空
				return 0
			}else {
				return -1
			}
		*/

	} else {
		//如果锁打开了，那么v1肯定是还有值的
		temp1, _ := strconv.Atoi(a.String())
		//如果V2结束了，那么肯定v1大
		if j == m {
			return 1
		}
		//继续将V2进行到底
		for ; j < m; j++ {
			if version2[j] == '.' {
				//这时候和上面的有所不同，逻辑不一样
				//这个时候v2找到了一个结束点
				temp2, _ := strconv.Atoi(b.String())

				if temp1 > temp2 {
					return 1
				} else if temp1 < temp2 {
					return -1
				}
				//如果这么比完之后，还是相等，继续搜索v2，只要
				for ; j < m; j++ {
					//换一个角度想，如果V2即不等于'.'又不等于'0',那么是肯定大于v1的
					if version2[j] == '.' || version2[j] == '0' {
						continue
					} else {
						return -1
					}
				}
				return 0

			} else if version2[j] == '0' {
				if mutex2 == false {
					continue
				} else {
					b.WriteByte(version2[j])
				}
			} else {
				mutex2 = true
				b.WriteByte(version2[j])
			}
		}
		if b.String() == "" {
			return 1
		}
		temp2, _ := strconv.Atoi(b.String())
		if temp1 > temp2 {
			return 1
		} else if temp1 < temp2 {
			return -1
		}
		return 0

	}
}

/*


func compareVersion(version1 string, version2 string) int {
	//解法应该是双指针
	n, m := len(version1), len(version2)
	i, j := 0, 0
	for ; i != n; i++ {
		if version1[i] == '.' {
			//为'.'的时候，需要找到一个有效数字
			i++
			for version1[i] == '0' || version1[i] == '.' {
				i++
				if i == n {
					break
				}
			}
			if i == n {
				break
			}
			//这样后version1[i]就为一个有效数字了
		}
		//然后为j寻找一个有效数字
		//需要先判断一个j是否超出长度，如果在i有有效数字的情况下，j超出长度，那么可以确定的是i>j
		if j == m {
			return 1
		}
		if version2[j] == '.' {
			//为'.'的时候，需要找到一个有效数字
			j++
			for version2[j] == '0' || version2[j] == '.' {
				j++
				if j == m {
					return 1
				}
			}
			if j == m {
				return 1
			}
			//这样后version2[j]就为一个有效数字了
		}

		//在都为有效数字的情况下进行比较
		if version1[i] > version2[j] {
			return 1
		} else if version1[i] < version2[j] {
			return -1
		}
		//如果相等，继续
		j++
	}

	//这时候，i的长度肯定结束了,j还有余粮，但是不代表j一定大于i
	if j == m {
		return 0
	} else {
		//for version2[j] == '0' || version2[j] == '.' {
		//	j++
		//	if j == m {
		//		return 0
		//	}
		//}
		return -1
	}
}
*/
