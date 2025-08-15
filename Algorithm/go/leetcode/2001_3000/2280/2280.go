package main

import "fmt"

func minimumLines(stockPrices [][]int) int {
	res := 1
	for i := 1; i < len(stockPrices); i++ {

	}

	return res
}

func main() {
	testCases := []struct {
		stockPrices [][]int
		expected    int
	}{
		{
			stockPrices: [][]int{{}, {}, {}},
			expected:    0,
		},
	}
	for i, tc := range testCases {
		res := minimumLines(tc.stockPrices)
		if res != tc.expected {
			fmt.Printf("Case #%d fail, expected %d got %d\n", i, tc.expected, res)
		} else {
			fmt.Printf("Case #%d pass", i)
		}
	}
}
