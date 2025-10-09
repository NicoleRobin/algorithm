package main

import (
	"fmt"
	"math"
)

/*
思路：动态规划

难点：倒序的动态规划
*/
func maximumEnergy(energy []int, k int) int {
	if len(energy) == 0 {
		return 0
	}

	ans := math.MinInt
	dp := make([]int, len(energy))
	for i := len(energy) - 1; i >= 0; i-- {
		if i+k <= len(energy)-1 {
			dp[i] = dp[i+k] + energy[i]
		} else {
			dp[i] = energy[i]
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		energy []int
		k      int
		expect int
	}{
		{
			energy: []int{5, 2, -10, -5, 1},
			k:      3,
			expect: 3,
		},
	}
	for _, testCase := range testCases {
		ans := maximumEnergy(testCase.energy, testCase.k)
		if ans != testCase.expect {
			fmt.Printf("\nExpected: %d, Actual: %d\n", testCase.expect, ans)
		} else {
			fmt.Printf("\nActual: %d\n", ans)
		}
	}
}
