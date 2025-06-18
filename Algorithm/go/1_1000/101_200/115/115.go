package main

/*
思路：动态规划
dp[i][j]表示s[i:]中有多少个t[j:]的子序列
*/
func numDistinct1(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}

	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
		dp[i][n] = 1
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][0]
}

/*
思路：动态规划
dp[i][j]表示t[:j]在s[:i]中的子序列个数
状态转移方程：
如果t[j] == s[i]:
	dp[i][j] = d[i-1][j-1] + dp[i-1][j]
否则:
	dp[i][j] = dp[i-1][j]
*/
func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}

	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
		// 初始值
		dp[i][0] = 1
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s[i] == t[j] {
				dp[i+1][j+1] = dp[i][j] + dp[i][j+1]
			} else {
				dp[i+1][j+1] = dp[i][j+1]
			}
		}
	}
	return dp[m][n]
}

func main() {
	testCases := []struct {
		s        string
		t        string
		expected int
	}{
		{
			s:        "rabbbit",
			t:        "rabbit",
			expected: 3,
		},
		{
			s:        "babgbag",
			t:        "bag",
			expected: 5,
		},
	}

	for i, tc := range testCases {
		ans := numDistinct(tc.s, tc.t)
		if ans != tc.expected {
			println("Test case", i, "failed: got", ans, "expected", tc.expected)
		} else {
			println("Test case", i, "passed")
		}
	}
}
