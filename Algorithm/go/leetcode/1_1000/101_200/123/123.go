package main

import "math"

/*
思路：动态规划
1、状态定义
dp[j][0]: 第j次交易后未持有股票时获得的利润
dp[j][1]: 第j次交易后持有股票时获得的利润
2、状态转移方程
3、初始状态

4、结果

*/
func maxProfit(prices []int) int {
	const k = 2
	f := [k + 2][2]int{}
	for j := 1; j <= k+1; j++ {
		f[j][1] = math.MinInt / 2 // 防止溢出
	}
	f[0][0] = math.MinInt / 2
	for _, p := range prices {
		for j := k + 1; j > 0; j-- {
			f[j][0] = max(f[j][0], f[j][1]+p)
			f[j][1] = max(f[j][1], f[j-1][0]-p)
		}
	}
	return f[k+1][0]
}

/*
思路：动态规划
该题是188题的特例，是188题中的k等于2时的情况；
*/
func maxProfit1(prices []int) int {
	const k = 2
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
		prices []int
		want   int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5, // 最大利润为6-1=5
		},
	}
	for i, tc := range testCases {
		ans := maxProfit(tc.prices)
		if ans != tc.want {
			println("Test case", i, "failed: expected", tc.want, "but got", ans)
		} else {
			println("Test case", i, "passed")
		}
	}
}
