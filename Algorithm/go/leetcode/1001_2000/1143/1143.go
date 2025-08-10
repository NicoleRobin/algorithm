package main

/*
思路：动态规划
dp[i][j]: 表示text1的前i个字符和text2的前j个字符的最长公共子序列长度
状态转移方程：
dp[i][j] = dp[i-1][j-1] + 1, if text1[i-1] == text2[j-1]
dp[i][j] = max(dp[i-1][j], dp[i][j-1]), if text1[i-1] != text2[j-1]
*/
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

type TestCase struct {
	text1 string
	text2 string
}

func main() {
	testcases := []TestCase{
		{"abcde", "ace"},
		{"abc", "abc"},
	}

	for _, testCase := range testcases {
		result := longestCommonSubsequence(testCase.text1, testCase.text2)
		println(result)
	}
}
