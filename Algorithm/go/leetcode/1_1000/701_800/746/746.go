package main

/*
dp[i] = min(dp[i-1]+cost[i-1], dp[i-2] + cost[i-2])
*/
func minCostClimbingStairs(cost []int) int {
	if len(cost) < 2 {
		return 0
	}

	dp := make([]int, len(cost)+1)
	for i := 2; i < len(dp); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(dp)-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	cost := []int{10, 15, 20}
	result := minCostClimbingStairs(cost)
	println(result) // Output: 15
}
