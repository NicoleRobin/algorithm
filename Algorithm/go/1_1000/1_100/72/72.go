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

type TestCase struct {
	word1 string
	word2 string
}

func main() {
	testCases := []TestCase{
		// {"horse", "ros"},
		{"pneumonoultramicroscopicsilicovolcanoconiosis", "ultramicroscopically"},
	}

	for _, testCase := range testCases {
		result := minDistance(testCase.word1, testCase.word2)
		println(result)
	}
}
