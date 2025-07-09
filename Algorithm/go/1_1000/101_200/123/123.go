package main

import "math"

/*
思路：动态规划
1、
*/
func maxProfit(prices []int) int {
	const k = 2
	f := [k + 2][2]int{}
	for j := 1; j <= k+1; j++ {
		f[j][1] = math.MinInt / 2 // 防止溢出
	}
	f[0][0] = math.MinInt / 2
	for _, p := range prices {
		for j := k + 1; j > 0; j-- {
			f[j][0] = max(f[j][0], f[j][1]+p)
			f[j][1] = max(f[j][1], f[j-1][0]-p)
		}
	}
	return f[k+1][0]
}

func main() {
	testCases := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5, // 最大利润为6-1=5
		},
	}
	for i, tc := range testCases {
		ans := maxProfit(tc.prices)
		if ans != tc.want {
			println("Test case", i, "failed: expected", tc.want, "but got", ans)
		} else {
			println("Test case", i, "passed")
		}
	}
}
