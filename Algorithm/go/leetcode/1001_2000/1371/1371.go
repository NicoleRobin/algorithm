package main

import "fmt"

/*
思考：
1、5个元音字母都必须出现么？
不是
2、出现的元音字母必须在这个子字符串中么？
不是

思路：
1、用一个数组记录从开头到当前位置的元音字母出现的次数
2、然后遍历前缀和数组，并判断之间的每个元音字母个数差值是否均为偶数

时间复杂度：O(n^2)
空间复杂度：O(n)
*/
func findTheLongestSubstring(s string) int {
	preSum := make([]map[byte]int, len(s)+1)
	preSum[0] = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		preSum[i+1] = make(map[byte]int)
		for _, v := range []byte{'a', 'e', 'i', 'o', 'u'} {
			if s[i] == v {
				preSum[i+1][v] = preSum[i][v] + 1
			} else {
				preSum[i+1][v] = preSum[i][v]
			}
		}
	}
	for i := 0; i < len(preSum); i++ {
		fmt.Println("i:", i, "preSum[i]:", preSum[i])
	}

	var ans int
	for i := 0; i < len(preSum); i++ {
		for j := i + 1; j < len(preSum); j++ {
			allEven := true
			for _, v := range []byte{'a', 'e', 'i', 'o', 'u'} {
				if (preSum[j][v]-preSum[i][v])%2 != 0 {
					allEven = false
					break
				}
			}

			if allEven {
				if j-i > ans {
					ans = j - i
				}
			}
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		s        string
		expected int
	}{
		{
			s:        "eleetminicoworoep",
			expected: 13,
		},
	}

	for i, tc := range testCases {
		result := findTheLongestSubstring(tc.s)
		if result != tc.expected {
			fmt.Println("Test case", i+1, "failed: expected", tc.expected, "but got", result)
		} else {
			fmt.Println("Test case", i+1, "passed")
		}
	}
}
