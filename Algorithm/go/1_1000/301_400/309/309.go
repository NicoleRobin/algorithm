package main

import (
	"fmt"
	"math"
)

/*
思路：动态规划
状态定义：dp[i][0]表示第i天持有股票的最大收益，dp[i][1]表示第i天不持有股票的最大收益
状态转移方程：
dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i]) // 第i天持有股票，可以是前一天持有，或者前一天不持有并买入
dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]) // 第i天不持有股票，可以是前一天不持有，或者前一天持有并卖出
*/
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n+2)
	dp[1][1] = math.MinInt
	for i, p := range prices {
		dp[i+2][0] = max(dp[i+1][0], dp[i+1][1]+p) // 持有股票
		dp[i+2][1] = max(dp[i+1][1], dp[i+1][0]-p) // 不持有股票
	}
	return dp[n+1][1] // 返回最后一天不持有股票的最大收益
}

func main() {
	testCases := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{1, 2, 3, 0, 2},
			want:   3,
		},
		{
			prices: []int{1, 2},
			want:   1,
		},
	}
	for i, tc := range testCases {
		result := maxProfit(tc.prices)
		if result == tc.want {
			fmt.Println("Test case", i+1, "passed.")
		} else {
			fmt.Println("Test case", i+1, "failed. Expected:", tc.want, "Got:", result)
		}
	}
}
