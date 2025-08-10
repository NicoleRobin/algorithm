package main

func minDistance(word1 string, word2 string) int {
	mLen, nLen := len(word1), len(word2)
	if mLen*nLen == 0 {
		return mLen + nLen
	}

	dp := make([][]int, mLen+1)
	for i := range dp {
		dp[i] = make([]int, nLen+1)
	}
	// 边界状态初始化
	for i := 0; i < mLen+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < nLen+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < mLen+1; i++ {
		for j := 1; j < nLen+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(min(dp[i][j-1]+1, dp[i-1][j]+1), dp[i-1][j-1])
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 2025-06-16
/*
思路：动态规划
dp[i][j]: word1的前i个字符和word2的前j个字符的最小编辑距离
*/
func minDistance1(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// 处理边界情况
	for i := 1; i < m+1; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}
	for j := 1; j < n+1; j++ {
		dp[0][j] = dp[0][j-1] + 1
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				// 如果相等，则不需要操作，继承上一个状态
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 如果不相等，则需要进行插入、删除或替换操作
				dp[i][j] = min(min(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1]+1)
			}
		}
	}

	return dp[m][n]
}

func main() {
	type TestCase struct {
		word1 string
		word2 string
	}
	testCases := []struct {
		word1, word2 string
		expect       int
	}{
		{
			word1:  "horse",
			word2:  "ros",
			expect: 3,
		},
		{
			word1:  "",
			word2:  "a",
			expect: 1,
		},
	}

	for i, tc := range testCases {
		ans := minDistance1(tc.word1, tc.word2)
		if ans != tc.expect {
			println("Test case", i, "failed: got", ans, "expected", tc.expect)
		} else {
			println("Test case", i, "passed")
		}
	}
}
