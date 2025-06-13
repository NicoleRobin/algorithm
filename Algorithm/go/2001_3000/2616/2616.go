package main

import (
	"fmt"
	"sort"
)

/*
思路：动态规划
dp[i]表示前i个数的最小最大值
dp[i] = min(dp[i-1], dp[i-2] + 1)
*/
func minimizeMax1(nums []int, p int) int {
	dp := make([]int, len(nums)+2)
	for i := 0; i < p; i++ {
		dp[i+2] = max(dp[i]+1, dp[i+1])
	}

	return dp[len(nums)+1]
}

func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)

	check := func(mx int) bool {
		cnt := 0
		for i := 0; i < len(nums)-1; i++ {
			if nums[i+1]-nums[i] <= mx {
				cnt++
				i++ // 跳过下一个数
			}
		}
		return cnt >= p
	}

	left, right := 0, nums[len(nums)-1]-nums[0]
	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid // 尝试更小的最大值
		} else {
			left = mid + 1 // 增大最大值
		}
	}

	return left
}

func main() {
	testCases := []struct {
		nums     []int
		p        int
		expected int
	}{
		{
			nums:     []int{10, 1, 2, 7, 1, 3},
			p:        2,
			expected: 1,
		},
		{
			nums:     []int{4, 2, 1, 2},
			p:        1,
			expected: 0,
		},
		{
			nums:     []int{1, 1, 0, 3},
			p:        2,
			expected: 2,
		},
	}

	for i, testCase := range testCases {
		result := minimizeMax(testCase.nums, testCase.p)
		fmt.Println("i:", i, "result:", result, "expected:", testCase.expected)
	}
}
