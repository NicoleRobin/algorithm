package main

import (
	"fmt"
	"math"
)

/*
思路：动态规划
*/
func numberOfWays(n int, x int) int {
	const mod = 1_000_000_007
	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, n+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		val := int64(math.Pow(float64(i), float64(x)))
		for j := 0; j <= n; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= int(val) {
				dp[i][j] = (dp[i][j] + dp[i-1][j-int(val)]) % mod
			}
		}
	}
	return int(dp[n][n])
}

func main() {
	testCases := []struct {
		n        int
		x        int
		expected int
	}{
		{
			n:        10,
			x:        2,
			expected: 1,
		},
		{
			n:        4,
			x:        1,
			expected: 2,
		},
	}
	for i, tc := range testCases {
		res := numberOfWays(tc.n, tc.x)
		if res != tc.expected {
			fmt.Printf("Case %d fail, res:%d, expected:%d\n", i, res, tc.expected)
		} else {
			fmt.Printf("Case %d pass\n", i)
		}
	}
}
