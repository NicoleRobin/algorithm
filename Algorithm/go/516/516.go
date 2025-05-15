package main

import (
	"fmt"
)

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
