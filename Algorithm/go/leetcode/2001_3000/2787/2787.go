package main

import (
	"fmt"
	"math"
)

/*
思路：动态规划
1、状态定义：dp[i]: i的方案数
2、状态转移方程：dp[i] = dp[i-1^x] + dp[i-2^x] + ...
*/
func numberOfWays(n int, x int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		var tmpSum int
		for j := 1; j <= i; j++ {
			powerOfJ := int(math.Pow(float64(j), float64(x)))
			if powerOfJ >= i {
				break
			}
			tmpSum += dp[i-powerOfJ]
		}
		dp[i] = tmpSum
	}
	fmt.Println("dp:", dp)

	return dp[n]
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
