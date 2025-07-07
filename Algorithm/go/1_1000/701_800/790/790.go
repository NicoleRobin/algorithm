package main

func numTilings(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 5
	for i := 4; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + 2 + dp[i-3] + 5
	}
	return dp[n]
}

func main() {

}
