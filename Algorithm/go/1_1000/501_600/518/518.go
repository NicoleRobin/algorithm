package main

/*
思考：
1、首先并没有想到为什么这道题适合动态规划，但是因为是动态规划的主体，所以。。。

思路：动态规划
状态定义：dp[i][j]表示使用前i种硬币凑成金额j的组合数
状态转移方程：dp[i][j] = dp[i-1][j] + dp[i][j-coin[i-1]]
初始状态：dp[0][0] = 1，表示使用0种硬币凑成金额0的组合数为1
结果：dp[len(coin)][amount]，表示使用所有硬币凑成金额amount的组合数
*/
func change(amount int, coins []int) int {
	var ans int
	n := len(coins)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}
	for i, coin := range coins {
		for i := 0; i <= amount; i++ {
		}
	}

	return dp[n][amount]
}

func main() {

}
