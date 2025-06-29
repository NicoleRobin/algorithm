package main

import (
	"fmt"
	"math"
)

/*
思考：
1、这道题特殊点是有冷冻期，所以需要考虑这个点对于状态转移方程的影响；

思路：动态规划
状态定义：
dp[i][0]: 表示第i天持有股票的最大收益
dp[i][1]: 表示第i天不持有股票的最大收益
状态转移方程：
dp[i][0] = max(dp[i-1][0], dp[i-2][1] - prices[i]) // 前一天持有股票，或者前两天不持有股票并且今天买入
dp[i][1] = max(dp[i-1][1], dp[i-1][0] + prices[i]) // 前一天不持有股票，或者前两天持有股票并且今天卖出
初始状态：
dp[0][0] = -prices[0] // 第一天买入股票
dp[0][1] = 0 // 第一天不持有股票

总结：
1、还是官方的题解，将状态定义为3个状态更好理解，dp[i][0]表示持有股票的收益；dp[i][1]表示不持有股票且处于冷冻期的收益；dp[i][2]表示不持有股票且不处于冷冻期的收益。
*/
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n+2)
	// 为什么要设置为math.MinInt？
	// 对于第 1 天 (i=1)，dp[i-2][1] 是无效的，因为没有第 -1 天的状态。因此，为了避免非法状态影响结果，dp[1][0] 被设置为一个极小值 (math.MinInt)，表示不可能的收益。
	dp[1][0] = math.MinInt
	for i, price := range prices {
		dp[i+2][0] = max(dp[i+1][0], dp[i][1]-price)   // 持有股票
		dp[i+2][1] = max(dp[i+1][1], dp[i+1][0]+price) // 不持有股票
	}
	// 循环中i的最大值为n-1，所以需要加2来处理最后一天的状态
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
