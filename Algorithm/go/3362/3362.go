package main

import "fmt"

/*
思考：
1、怎么挑选该删除哪个？
2、前面两个题都是先计算查询数组能否覆盖的范围，然后判断是否大于等于对应位置的数值；那这次
是否能否反向思考？
3、
*/
func maxRemoval(nums []int, queries [][]int) int {
	deltaArray := make([]int, len(nums)+1)
	for _, query := range queries {
		left, right := query[0], query[1]
		deltaArray[left] += 1
		deltaArray[right+1] -= 1
	}

	var prefixSum int
	prefixSumArray := make([]int, len(nums)+1)
	for i, delta := range deltaArray {
		prefixSum += delta
		prefixSumArray[i] = prefixSum
	}
	fmt.Println("prefixSumArray:", prefixSumArray)

	diffArray := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		diffArray[i] = prefixSumArray[i] - nums[i]
	}

	return 0
}

func main() {
	testCases := []struct {
		nums    []int
		queries [][]int
		expect  int
	}{
		{
			nums:    []int{2, 0, 2},
			queries: [][]int{{0, 2}, {0, 2}, {1, 1}},
			expect:  1,
		},
	}
	for i, tc := range testCases {
		result := maxRemoval(tc.nums, tc.queries)
		if result != tc.expect {
			println("Test case", i, "failed: expected", tc.expect, "but got", result)
		} else {
			println("Test case", i, "passed")
		}
	}
}
