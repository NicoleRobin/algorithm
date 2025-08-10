package main

import "math"

/*
思路：动态规划
dp[i]定义凑成总数为i所需的最少硬币个数；
状态转移方程：
dp[i] = min(j=0..i-1)dp[j] + 1
*/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		min := math.MaxInt32
		for _, coin := range coins {
			if i-coin >= 0 {
				if dp[i-coin] != -1 && dp[i-coin] < min {
					min = dp[i-coin]
				}
			}
		}
		if min != math.MaxInt32 {
			dp[i] = min + 1
		} else {
			dp[i] = -1
		}
	}
	return dp[amount]
}

func main() {

}
