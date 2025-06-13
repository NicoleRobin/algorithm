package main

import "fmt"

/*
思路：动态规划
1、dp[i][j] 表示 s[i:j] 是否为回文串
2、状态转移方程：
dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
*/
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	var begin, maxLen int
	// 枚举子串长度
	for L := 2; L <= n; L++ {
		// 枚举左边界
		for i := 0; i < n; i++ {
			// 由L和i可以确定右边界，即j-i+1=L
			j := i + L - 1
			if j >= n {
				break
			}

			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					// 一个字符或者两个字符时
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && L > maxLen {
				maxLen = L
				begin = i
			}
		}
	}

	return s[begin : begin+maxLen]
}

func main() {
	testCases := []string{
		"babad",
	}
	for _, testCase := range testCases {
		result := longestPalindrome(testCase)
		fmt.Println("result:", result)
	}
}
