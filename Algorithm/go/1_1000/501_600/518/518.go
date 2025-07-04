package main

import "fmt"

/*
思考：
1、首先并没有想到为什么这道题适合动态规划，但是因为是动态规划的主体，所以。。。

思路：动态规划
状态定义：dp[i]表示金额之和等于x的硬币组合数
状态转移方程：dp[i] = dp[i-coin]{for coin in coins} + 1
初始状态：dp[0] = 1，表示凑成金额0的组合数为1
结果：dp[amount]，表示凑成金额amount的组合数
*/
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = dp[i-coin] + 1
			}
		}
	}

	return dp[amount]
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
