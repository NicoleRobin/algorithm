package main

import "fmt"

/*
思考：

思路：模拟题
1、
*/
func minZeroArray(nums []int, queries [][]int) int {
	ans := -1
	for i, query := range queries {
		var nonZeroCount int
		for _, num := range nums {
			if num != 0 {
				nonZeroCount++
			}
		}
		if nonZeroCount == 0 {
			ans = i
			break
		}

		for j := query[0]; j <= query[1]; j++ {
			if nums[j] < query[2] {
				nums[j] = 0
			} else {
				nums[j] -= query[2]
			}
		}

		nonZeroCount = 0
		for _, num := range nums {
			if num != 0 {
				nonZeroCount++
			}
		}
		if nonZeroCount == 0 {
			ans = i + 1
			break
		}
	}
	return ans
}

func main() {
	testCases := []struct {
		nums     []int
		queues   [][]int
		expected int
	}{
		{
			nums:     []int{2, 0, 2},
			queues:   [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}},
			expected: 2,
		},
		{
			nums:     []int{0},
			queues:   [][]int{{0, 0, 2}, {0, 0, 4}, {0, 0, 3}, {0, 0, 5}},
			expected: 0,
		},
	}

	for i, testCase := range testCases {
		result := minZeroArray(testCase.nums, testCase.queues)
		fmt.Println("i:", i, "result:", result, "expected:", testCase.expected)
	}
}
