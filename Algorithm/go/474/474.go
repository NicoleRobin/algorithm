package main

import (
	"strings"
)

/*
思路：动态规划-01背包问题
dp[i][j]表示：在最多i个0和j个1的情况下，最多能够选多少个字符串

*/
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	var ans int
	for _, str := range strs {
		zeroCount := strings.Count(str, "0")
		oneCount := len(str) - zeroCount
		for j := m; j >= zeroCount; j-- {
			for k := n; k >= oneCount; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zeroCount][k-oneCount]+1)
				ans = max(ans, dp[j][k])
			}
		}
	}

	return dp[m][n]
}

func main() {
}
