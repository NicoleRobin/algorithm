package main

import "fmt"

func maxProfit(prices []int, strategy []int, k int) int64 {
	presumPrices := make([]int, len(prices)+1)
	presums := make([]int, len(prices)+1)
	var presum, presum1 int
	for i := 0; i < len(prices); i++ {
		presum += prices[i] * strategy[i]
		presum1 += prices[i]
		presumPrices[i+1] = presum
		presums[i+1] = presum1
	}

	var postsum int
	postsumPrices := make([]int, len(prices))
	for i := len(prices) - 1; i >= 0; i-- {
		postsum += prices[i] * strategy[i]
		postsumPrices[i] = postsum
	}
	fmt.Println("presumPrices:", presumPrices)
	fmt.Println("postsumPrices:", postsumPrices)
	fmt.Println("presums:", presums)

	var res int64 = int64(presum)
	for left := 0; left+k-1 < len(prices); left++ {
		var profit int
		profit += presumPrices[left]
		if left+k < len(prices) {
			profit += postsumPrices[left+k]
		}
		profit += presums[left+k] - presums[left+k/2]
		if int64(profit) > res {
			res = int64(profit)
		}
	}
	return res
}

func main() {
	testCase := []struct {
		prices   []int
		strategy []int
		k        int
		expected int64
	}{
		{
			prices:   []int{4, 2, 8},
			strategy: []int{-1, 0, 1},
			k:        2,
			expected: 10,
		},
	}
	for i, tc := range testCase {
		res := maxProfit(tc.prices, tc.strategy, tc.k)
		if res != tc.expected {
			fmt.Printf("case #%d fail: %d %d\n", i+1, res, tc.expected)
		} else {
			fmt.Printf("case #%d pass: %d %d\n", i+1, res, tc.expected)
		}
	}
}
