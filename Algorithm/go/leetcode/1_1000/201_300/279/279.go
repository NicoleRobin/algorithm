package main

import "math"

/*
思路：动态规划
1、dp[i] = min(dp[i-j*j]) + 1
*/
func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		minAns := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			if dp[i-j*j] < minAns {
				minAns = dp[i-j*j]
			}
		}
		dp[i] = minAns + 1
	}
	return dp[n]
}

func main() {

}
