package main

import (
	"fmt"
	"sort"
)

/*
思路：动态规划
1、状态定义
dp[i]: 第i个信封作为结尾的最大信封个数
2、状态转移方程
dp[i] = 0到i-1中，信封的宽和高都小于第i个信封的最大值

3、初始状态

4、结果
dp数组的最大值
*/
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] < envelopes[j][1]
		} else {
			return false
		}
	})
	n := len(envelopes)
	dp := make([]int, n+1)
	dp[0] = 1

	var ans int
	for i := 0; i < n; i++ {
		maxValue := 0
		for j := 0; j < i; j++ {
			if envelopes[j][0] >= envelopes[i][0] {
				break
			}

			if envelopes[j][0] < envelopes[i][0] && envelopes[j][1] < envelopes[i][1] {
				if dp[j] > maxValue {
					maxValue = dp[j]
				}
			}
		}
		dp[i] = maxValue + 1

		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

func main() {
	testCases := []struct {
		envelopes [][]int
		want      int
	}{
		{
			envelopes: [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}},
			want:      3,
		},
		{
			envelopes: [][]int{{1, 1}, {1, 1}, {1, 1}},
			want:      1,
		},
		{
			envelopes: [][]int{{1, 1}},
			want:      1,
		},
	}

	for i, testCase := range testCases {
		ans := maxEnvelopes(testCase.envelopes)
		if ans != testCase.want {
			fmt.Printf("Case #%d: failed, ans=%d, want %d\n", i+1, ans, testCase.want)
		} else {
			fmt.Printf("Case #%d: passed\n", i+1)
		}
	}
}
