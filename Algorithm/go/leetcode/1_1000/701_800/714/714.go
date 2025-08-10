package main

/*
思考：
1、虽然题目不限制交易次数，但是每次交易都需要支付手续费，所以还是应该减少交易次数。
2、

思路：动态规划
状态定义：
dp[i][0]表示第 i 天交易完后手里没有股票的最大利润
dp[i][1]表示第 i 天交易完后手里持有一支股票的最大利润（i 从 0 开始）
状态转移方程：
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i] - fee)
dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
*/
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[n-1][0]
}

func main() {

}
