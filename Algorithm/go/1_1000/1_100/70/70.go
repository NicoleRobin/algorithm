package main

/*
思路：动态规划
1、dp[i]=max(dp[i−2]+nums[i],dp[i−1])
*/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	if nums[1] > nums[0] {
		dp[1] = nums[1]
	} else {
		dp[1] = nums[0]
	}

	var result int
	for i := 0; i < len(nums); i++ {
		if i >= 2 {
			if dp[i-2]+nums[i] > dp[i-1] {
				dp[i] = dp[i-2] + nums[i]
			} else {
				dp[i] = dp[i-1]
			}
		}

		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	nums := [][]int{{1, 2, 3, 1}, {2, 7, 9, 3, 1}}
	for _, num := range nums {
		println(rob(num))
	}
}
