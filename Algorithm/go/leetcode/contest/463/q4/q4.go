package main

import "fmt"

const (
	mod = 1000000007
)

/*
思考：
1、该题是第二题的优化版，测试用例数据量从10的3次方提升到10的5次方，意味着算法必须是O(n)才可以了；
2、
*/
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
