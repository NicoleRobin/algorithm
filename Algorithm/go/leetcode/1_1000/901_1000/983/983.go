package main

import "fmt"

/*
思考：
1、没办法直观的看出来状态的定义，假如定义状态为dp[i]:完成前i个计划旅行所需的最低消费，那它其实跟之前是买的什么票有关系（1天、7天、30天）；
2、如果把dp[i][j]定义为完成前i个计划旅行所需的最低消费呢？

思路：动态规划
1、定义状态：dp[i] 表示完成前i天旅行所需的最低消费
2、状态转移方程：

	dp[i] = min(dp[i-1]+costs[0], dp[max(i-7, 0)]+costs[1], dp[max(i-30, 0)]+costs[2])

3、初始化：dp[0] = 0
4、边界条件：dp[i] = dp[i-1]
5、返回结果：dp[lastDay]
*/
func mincostTickets(days []int, costs []int) int {
	if len(days) == 0 {
		return 0
	}
	dayMap := make(map[int]bool)
	for _, day := range days {
		dayMap[day] = true
	}

	lastDay := days[len(days)-1]
	dp := make([]int, lastDay+1)
	for i := 1; i <= lastDay; i++ {
		if dayMap[i] {
			dp[i] = min(dp[i-1]+costs[0], dp[max(i-7, 0)]+costs[1], dp[max(i-30, 0)]+costs[2])
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[lastDay]
}

func main() {
	testCases := []struct {
		days  []int
		costs []int
		want  int
	}{
		{[]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}, 11},
	}

	for i, testCase := range testCases {
		ans := mincostTickets(testCase.days, testCase.costs)
		if ans != testCase.want {
			fmt.Printf("i: %d failed, ans: %v, want: %v\n", i, ans, testCase.want)
		} else {
			fmt.Printf("i: %d passed\n", i)
		}
	}
}
