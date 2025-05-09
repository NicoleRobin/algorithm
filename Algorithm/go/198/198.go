package main

/*
思路：动态规划
dp[i] = max(dp[i-1], dp[i-2]+nums[i])
*/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[len(dp)-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
}
