package main

import (
	"fmt"
	"strings"
)

/*
思路：动态规划
dp[i][j][k]: 前i个字符串中，j个0，k个1时，最大子集的长度；
状态转移方程：
dp[i][j][k] = max(dp[i-1][j][k], dp[i-zeroCount][j-oneCount]
*/
func findMaxForm1(strs []string, m int, n int) int {
	dp := make([][][]int, len(strs)+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}

	for i, str := range strs {
		cnt0 := strings.Count(str, "0")
		cnt1 := len(str) - cnt0
		for j := range m + 1 {
			for k := range n + 1 {
				if j >= cnt0 && k >= cnt1 {
					// 当可以选择第i个字符串时
					dp[i+1][j][k] = max(dp[i][j][k], dp[i][j-cnt0][k-cnt1]+1)
				} else {
					dp[i+1][j][k] = dp[i][j][k]
				}
			}
		}
	}

	return dp[len(strs)][m][n]
}

/*
空间优化
*/
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		zeroCount := strings.Count(str, "0")
		oneCount := len(str) - zeroCount
		for j := m; j >= zeroCount; j-- {
			for k := n; k >= oneCount; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zeroCount][k-oneCount]+1)
			}
		}
	}

	return dp[m][n]
}

func main() {
	testCases := []struct {
		strs   []string
		m, n   int
		expect int
	}{
		{
			strs: []string{},
			m:    5,
			n:    3,
		},
	}
	for i, tc := range testCases {
		ans := findMaxForm(tc.strs, tc.m, tc.n)
		fmt.Println("i:", i, "ans:", ans, "expect:", tc.expect)
	}
}
