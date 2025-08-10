package main

import "fmt"

/*
思考：
1、和之前的凑硬币问题518一样；
2、看过答案后发现和518并不相同，但是为啥呢？

思路：动态规划
1、
*/
func combinationSum41(nums []int, target int) int {
	n := len(nums)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 1

	for i, num := range nums {
		for c := 0; c <= target; c++ {
			if c < num {
				dp[i+1][c] = dp[i][c]
			} else {
				dp[i+1][c] = dp[i+1][c-num] + dp[i][c]
			}
		}
	}
	fmt.Println(dp)

	return dp[n][target]
}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

func main() {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{1, 2, 3},
			target: 4,
			want:   7,
		},
	}
	for i, tc := range testCases {
		ans := combinationSum4(tc.nums, tc.target)
		if ans != tc.want {
			fmt.Println("test case ", i, " failed, want:", tc.want, " ans:", ans)
		} else {
			fmt.Println("test case ", i, " pass")
		}
	}
}
