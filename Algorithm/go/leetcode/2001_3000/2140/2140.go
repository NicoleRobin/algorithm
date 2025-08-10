package main

import (
	"fmt"
	"slices"
)

/*
思考：
1、有点像是背包问题的变形，可以往背包问题解题套路思考；
2、分为选或者必选两种情况思考；

思路：动态规划
状态定义
dp[i][0]: 前i个节点中，不选第i个节点时获取的最高分
dp[i][1]: 前i个节点中，选择第i个节点时获取的最高分
状态转移方程：
dp[i][0] = max(dp[i-1][0], dp[i-1][1])
dp[i][1] = points[i] + max(之前所有冷冻期已结束的状态)

反思：
这个思路定义的状态不遵循无后效性的要求，例如在考虑第i个节点时，选第i个节点时，需要从i前面所有冷冻期已结束或者没有选择某个节点的结果中选取最大值，
但是，这里有一个问题：假如前面有一个节点j，没有选取节点i时的结果是最大值，但是有可能它的结果是从前面某一个有冷冻期的节点转变过来的，
而那个节点的冷冻期还没有结束，此时如果选择节点j的结果作为最大值并加上节点i的point，就有问题了；
*/
func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([][2]int64, n+1)
	for i := 0; i < len(questions); i++ {
		point := int64(questions[i][0])
		// fmt.Println("i:", i, " point:", point)

		dp[i+1][0] = max(dp[i][0], dp[i][1])
		var maxVal int64 = 0
		for j := 0; j < i; j++ {
			if j+questions[j][1] >= i {
				continue
			}

			if dp[j+1][0] > maxVal {
				maxVal = dp[j+1][0]
			}

			if dp[j+1][1] > maxVal {
				maxVal = dp[j+1][1]
			}
		}
		dp[i+1][1] = maxVal + point
		fmt.Println("i:", i, " dp:", dp)
	}

	return max(dp[n][0], dp[n][1])
}

func mostPoints1(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n+1)

	for i := n - 1; i >= 0; i-- {
		points := int64(questions[i][0])
		skip := questions[i][1]

		// 如果答了第i题，跳过skip个题
		next := i + skip + 1
		if next > n {
			dp[i] = max(dp[i+1], points)
		} else {
			dp[i] = max(dp[i+1], points+dp[next])
		}
	}

	return dp[0]
}

func mostPoints2(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n+1)
	for i, q := range slices.Backward(questions) {
		j := min(i+q[1]+1, n)
		dp[i] = max(dp[i+1], dp[j]+int64(q[0]))
	}
	return dp[0]
}

func main() {
	testCases := []struct {
		questions [][]int
		want      int64
	}{
		{
			questions: [][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}},
			want:      5,
		},
	}
	for i, tc := range testCases {
		ans := mostPoints(tc.questions)
		if tc.want != ans {
			fmt.Println("i:", i, " failed, ans:", ans, " want:", tc.want)
		} else {
			fmt.Println("i:", i, " pass")
		}
	}
}
