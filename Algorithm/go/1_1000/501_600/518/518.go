package main

import "fmt"

/*
思考：
1、首先并没有想到为什么这道题适合动态规划，但是因为是动态规划的主体，所以。。。

思路：动态规划
状态定义：dp[i][c]表示前i个金币中选择金额之和等于c的硬币组合数
状态转移方程：
当第i个硬币大于金额之和时：
	dp[i][c] = dp[i-1][c]
否则：
	dp[i][c] = dp[i][c-coin] + dp[i-1][c]
初始状态：dp[0][0] = 1，表示使用0个金币凑成金额0的组合数为1
结果：dp[n][amount]，表示使用所有金币凑成金额amount的组合数
*/
func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}
	dp[0][0] = 1

	for i, coin := range coins {
		for c := 0; c <= amount; c++ {
			if c < coin {
				dp[i+1][c] = dp[i][c]
			} else {
				dp[i+1][c] = dp[i+1][c-coin] + dp[i][c]
			}
		}
	}

	return dp[n][amount]
}

func main() {
	testCases := []struct {
		amount int
		coins  []int
		want   int
	}{
		{
			amount: 5,
			coins:  []int{1, 2, 5},
			want:   4,
		},
	}
	for i, tc := range testCases {
		ans := change(tc.amount, tc.coins)
		if ans != tc.want {
			fmt.Println("test case ", i, " failed, want:", tc.want, " ans:", ans)
		} else {
			fmt.Println("test case ", i, " pass")
		}
	}

}
