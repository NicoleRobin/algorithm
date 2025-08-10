package main

import (
	"fmt"
	"math"
)

/*
思路：动态规划
dp[i][0]表示第i天不持有股票的最大利润
dp[i][1]表示第i天持有股票的最大利润

状态转移方程：
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
*/
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n+1)
	// 初始化
	dp[0][1] = math.MinInt
	for i, price := range prices {
		dp[i+1][0] = max(dp[i][0], dp[i][1]+price) // 不持有股票
		dp[i+1][1] = max(dp[i][1], dp[i][0]-price) // 持有股票
	}

	return dp[n][0] // 最后一天不持有股票的最大利润
}

func main() {
	testCases := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   7,
		},
	}
	for i, tc := range testCases {
		ans := maxProfit(tc.prices)
		if ans != tc.want {
			fmt.Println("Test case", i, "failed: expected", tc.want, "but got", ans)
		} else {
			fmt.Println("Test case", i, "passed")
		}
	}
}
