package main

import "fmt"

/*
思路：动态规划
1、状态转移方程如下：
如果 j≥nums[i]，则对于当前的数字 nums[i]，可以选取也可以不选取，两种情况只要有一个为 true，就有 dp[i][j]=true。

如果不选取 nums[i]，则 dp[i][j]=dp[i−1][j]；
如果选取 nums[i]，则 dp[i][j]=dp[i−1][j−nums[i]]。
如果 j<nums[i]，则在选取的数字的和等于 j 的情况下无法选取当前的数字 nums[i]，因此有 dp[i][j]=dp[i−1][j]。
*/

func canPartition(nums []int) bool {
	var maxNum, sum int
	for _, num := range nums {
		sum += num
		if num > maxNum {
			maxNum = num
		}
	}

	if sum%2 != 0 {
		return false
	}
	halfSum := sum / 2
	if maxNum > halfSum {
		return false
	}

	dp := make([][]bool, len(nums))
	for i := range dp {
		dp[i] = make([]bool, halfSum+1)
	}
	for i := 0; i < len(nums); i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		for j := 1; j <= halfSum; j++ {
			if j >= v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(nums)-1][halfSum]
}

func canPartition1(nums []int) bool {
	var maxNum, sum int
	for _, num := range nums {
		sum += num
		if num > maxNum {
			maxNum = num
		}
	}

	if sum%2 != 0 {
		return false
	}
	halfSum := sum / 2
	if maxNum > halfSum {
		return false
	}

	dp := make([]bool, halfSum+1)
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		for j := halfSum; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}

	return dp[halfSum]
}

func main() {
	testCases := [][]int{
		// {1, 5, 11, 5},
		{3, 3, 3, 4, 5},
	}
	for _, testCase := range testCases {
		result := canPartition(testCase)
		fmt.Println("nums:", testCase, ",result:", result)
	}
}
