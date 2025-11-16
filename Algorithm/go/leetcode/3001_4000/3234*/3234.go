package main

import "fmt"

/*
思路：
1、求出所有子字符串
2、抽象一个函数判断一个字符串是否是1显著的字符串

超时了，想想如何优化
*/
func numberOfSubstrings(s string) int {
	var ans int
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			subS := s[j : i+1]
			fmt.Println("i:", i, ", j:", j, "subS:", subS)
			if isSpecialOne(subS) {
				ans++
			}
		}
	}

	return ans
}

func isSpecialOne(s string) bool {
	var oneCount int
	for _, ch := range s {
		if ch == '1' {
			oneCount++
		}
	}
	zeroCount := len(s) - oneCount
	return oneCount >= zeroCount*zeroCount
}

/*
思考：
1、可以考虑滑动窗口，当不符合时就移动右边的指针扩大窗口，直到符合要求；

问题：
这里和子数组和的问题有些差异，并不是移动右指针就会增大1的个数，也不是移动左指针就会减小0的个数，所以这样处理滑动操作就不对；
*/
func numberOfSubstrings2(s string) int {
	var ans int
	var left, right = 0, 0
	for right < len(s) {
		if !isSpecialOne(s[left : right+1]) {
			// 分两种情况：
		} else {
			for left <= right && isSpecialOne(s[left:right+1]) {
				ans++
				left++
			}
		}
	}

	return ans
}

/*
官方题解
1、枚举0的个数，因为1显著字符串的特性是1的个数大于等于0的个数的平方，所以有效的0个数是固定；
*/
func numberOfSubstrings3(s string) int {
	n := len(s)
	// 左边最近一个0的位置数组
	pre := make([]int, n+1)
	pre[0] = -1
	for i := 0; i < n; i++ {
		if i == 0 || (i > 0 && s[i-1] == '0') {
			pre[i+1] = i
		} else {
			pre[i+1] = pre[i]
		}
	}

	res := 0
	for i := 1; i <= n; i++ {
		cnt0 := 0
		if s[i-1] == '0' {
			cnt0 = 1
		}
		j := i
		// 枚举0的个数
		for j > 0 && cnt0*cnt0 <= n {
			cnt1 := (i - pre[j]) - cnt0
			if cnt0*cnt0 <= cnt1 {
				add := j - pre[j]
				if cnt1-cnt0*cnt0+1 < add {
					add = cnt1 - cnt0*cnt0 + 1
				}
				res += add
			}
			j = pre[j]
			cnt0++
		}
	}
	return res
}

func main() {
	testCases := []struct {
		s        string
		expected int
	}{
		{
			s:        "00011",
			expected: 5,
		},
	}
	for i, tc := range testCases {
		result := numberOfSubstrings3(tc.s)
		if result != tc.expected {
			fmt.Printf("Case: %d, expected:%d, result:%d\n", i, tc.expected, result)
		} else {
			fmt.Printf("Case: %d, pass\n", i)
		}
	}
}
