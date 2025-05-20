package main

import (
	"fmt"
)

/*
思路：动态规划
dp[i][j]表示s[i:j]的最长回文子序列长度

`i` 从 `n-1` 到 `0` 进行倒序遍历的原因是动态规划的状态转移需要依赖于已经计算出的子问题的结果。在这个算法中，`dp[i][j]` 的值依赖于 `dp[i+1][j-1]`、`dp[i+1][j]` 和 `dp[i][j-1]`。
如果 `i` 从小到大正序遍历，那么在计算 `dp[i][j]` 时，`dp[i+1][j-1]` 和 `dp[i+1][j]` 可能还没有被计算出来，导致无法正确地进行状态转移。因此，必须从 `n-1` 开始倒序遍历 `i`，确保在计算 `dp[i][j]` 时，所有依赖的状态都已经被计算完成。
这种倒序遍历的方式是动态规划中常见的技巧，尤其是在二维数组中需要依赖右下方或下方的状态时。
*/
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

func main() {
	testCases := []struct {
		s        string
		expected int
	}{
		{
			s:        "bbbab",
			expected: 4,
		},
		{
			s:        "cbbd",
			expected: 2,
		},
	}

	for _, testCase := range testCases {
		result := longestPalindromeSubseq(testCase.s)
		fmt.Println("Input:", testCase.s, "Expected:", testCase.expected, "Result:", result)
	}
}
