package main

import (
	"fmt"
	"math"
)

/*
思路：动态规划
1、状态定义
dp[i][j][0]: 到第i天结束时完成至多j笔交易，未持有股票的最大利润；
dp[i][j][1]: 到第i天结束时完成至多j笔交易，持有股票的最大利润；
2、状态转移方程
卖出
dp[i][j][0] = dp[i-1][j][1] + prices[i]
买入
dp[i][j][1] = dp[i-1][j-1][0] - prices[i]
什么都不做
dp[i][j][0] = dp[i-1][j][0]
dp[i][j][1] = dp[i-1][j][1]
3、初始状态

4、结果
因为最后持有股票肯定不能比未持有股票获得的利润更多，所以结果肯定是未持有股票的状态；
dp[n][k][0]
*/
func maxProfit(k int, prices []int) int {
	n := len(prices)
	dp := make([][][2]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][2]int, k+2)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j][0], dp[i][j][1] = math.MinInt32, math.MinInt32
		}
	}
	for j := 1; j < k+2; j++ {
		dp[0][j][0] = 0
	}

	for i := 0; i < n; i++ {
		for j := 1; j < k+2; j++ {
			dp[i+1][j][0] = max(dp[i][j][0], dp[i][j-1][1]+prices[i])
			dp[i+1][j][1] = max(dp[i][j][1], dp[i][j][0]-prices[i])
		}
	}

	return dp[n][k+1][0]
}

func main() {
	testCases := []struct {
		k      int
		prices []int
		want   int
	}{
		{2, []int{2, 4, 1}, 2},
		{2, []int{3, 2, 6, 5, 0, 3}, 7},
	}

	for i, testCase := range testCases {
		ans := maxProfit(testCase.k, testCase.prices)
		if ans != testCase.want {
			fmt.Printf("Case #%d: failed, ans=%d, want %d\n", i+1, ans, testCase.want)
		} else {
			fmt.Printf("Case #%d: passed\n", i+1)
		}
	}
}
