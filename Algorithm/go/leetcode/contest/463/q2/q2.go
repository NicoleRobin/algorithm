package main

import (
	"fmt"
)

const (
	mod = 1000000007
)

func xorAfterQueries(nums []int, queries [][]int) int {
	if len(nums) == 0 {
		return 0
	}

	for i := 0; i < len(queries); i++ {
		for j := queries[i][0]; j <= queries[i][1]; j += queries[i][2] {
			nums[j] = (nums[j] * queries[i][3]) % mod
		}
		fmt.Println(nums)
	}

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = res ^ nums[i]
	}

	return res
}

func main() {
	testCase := []struct {
		nums     []int
		queries  [][]int
		expected int
	}{
		{
			nums:     []int{780},
			queries:  [][]int{{0, 0, 1, 13}, {0, 0, 1, 17}, {0, 0, 1, 9}, {0, 0, 1, 18}, {0, 0, 1, 16}, {0, 0, 1, 6}, {0, 0, 1, 4}, {0, 0, 1, 11}, {0, 0, 1, 7}, {0, 0, 1, 18}, {0, 0, 1, 8}, {0, 0, 1, 15}, {0, 0, 1, 12}},
			expected: 0,
		},
		{
			nums:     []int{2, 3, 1, 5, 4},
			queries:  [][]int{{1, 4, 2, 3}, {0, 2, 1, 2}},
			expected: 31,
		},
	}
	for i, tc := range testCase {
		res := xorAfterQueries(tc.nums, tc.queries)
		if res != tc.expected {
			fmt.Printf("case #%d fail: %d %d\n", i+1, res, tc.expected)
		} else {
			fmt.Printf("case #%d pass: %d %d\n", i+1, res, tc.expected)
		}
	}
}
