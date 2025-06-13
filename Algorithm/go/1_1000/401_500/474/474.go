package main

import (
	"strings"
)

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
思路：动态规划

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
