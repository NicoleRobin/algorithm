package main

/*
思路：动态规划
1、先计算字符串 s 的最长回文子序列长度 spa。
2、最少插入次数等于字符串长度减去最长回文子序列长度，即 len(s) - spa。
*/
func minInsertions(s string) int {
	spa := longestPalindromeSubseq(s)
	return len(s) - spa
}

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

/*
思路：区间动态规划
1、状态定义：
dp[i][j]: 表示使字符串 s[i:j] 成为回文串所需的最少插入次数
2、
*/
func minInsertions2(s string) int {
	return 0
}

func main() {

}
