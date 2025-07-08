package main

import (
	"fmt"
)

/*
思考：
1、动态规划有两种思路：枚举、选或者不选，这道题应该是选或者不选类型的；

思路：动态规划
1、定义状态：dp[i] 表示长度为i的good string的个数
2、状态转移方程：

	dp[i] = dp[i-zero] + dp[i-one]

3、初始化：dp[0] = 1
4、返回结果：dp[low] + dp[low+1] ... dp[high]
*/
func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high+1)
	dp[0] = 1
	for i := 1; i <= high; i++ {
		if i >= zero {
			dp[i] += dp[i-zero]
		}
		if i >= one {
			dp[i] += dp[i-one]
		}
		dp[i] = dp[i] % (1e9 + 7)
	}

	var ans int
	for i := low; i <= high; i++ {
		ans += dp[i]
		ans = ans % (1e9 + 7)
	}
	return ans
}

func main() {
	testCases := []struct {
		low, high, zero, one int
		want                 int
	}{
		{
			low:  3,
			high: 3,
			zero: 1,
			one:  1,
			want: 8,
		},
	}
	for i, tc := range testCases {
		ans := countGoodStrings(tc.low, tc.high, tc.zero, tc.one)
		if ans != tc.want {
			fmt.Printf("Case #%d failed, want:%d, ans:%d\n", i+1, tc.want, ans)
		} else {
			fmt.Printf("Case #%d passwd\n", i+1)
		}
	}
}
