package main

/*
思路：贪心算法
1、遍历价格数组，记录当前的最低价格。
2、对于每一天的价格，计算当前价格与最低价格的差值，如果这个差值大于当前已知的最大利润，则更新最大利润。
*/
func maxProfit(prices []int) int {
	var ans int
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i] // 更新最低价格
		}
		if prices[i]-minPrice > ans {
			ans = prices[i] - minPrice // 更新最大利润
		}
	}

	return ans
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
