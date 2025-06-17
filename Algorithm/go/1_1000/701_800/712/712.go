package main

import "fmt"

func minimumDeleteSum1(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		if i > 0 {
			dp[i][0] = dp[i-1][0] + int(s1[i-1])
		}
	}
	for j := range dp[0] {
		if j > 0 {
			dp[0][j] = dp[0][j-1] + int(s2[j-1])
		}
	}
	for i, c1 := range s1 {
		for j, c2 := range s2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i][j+1]+int(c1), dp[i+1][j]+int(c2))
			}
		}
	}
	return dp[m][n]
}

// 2025-06-16
/*
思路：动态规划
dp[i][j]:s1的前i个字符和s2的前j个字符的最小删除和

类似题目：
LeetCode-72：https://leetcode.cn/problems/edit-distance/description/
LeetCode-583：https://leetcode.cn/problems/delete-operation-for-two-strings/description/
*/

func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		dp[i] = make([]int, len(s2)+1)
	}
	for i := 1; i < len(s1)+1; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}
	for j := 1; j < len(s2)+1; j++ {
		dp[0][j] = dp[0][j-1] + int(s2[j-1])
	}

	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+int(s1[i-1]), min(dp[i][j-1]+int(s2[j-1]), dp[i-1][j-1]+int(s1[i-1])+int(s2[j-1])))
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

func main() {
	testCases := []struct {
		s1, s2 string
		expect int
	}{
		{
			s1:     "sea",
			s2:     "eat",
			expect: 231,
		},
		{
			s1:     "delete",
			s2:     "leet",
			expect: 403,
		},
	}
	for i, tc := range testCases {
		ans := minimumDeleteSum(tc.s1, tc.s2)
		if ans != tc.expect {
			fmt.Println("Test case", i, "failed: expected", tc.expect, "but got", ans)
		} else {
			fmt.Println("Test case", i, "passed")
		}
	}
}
