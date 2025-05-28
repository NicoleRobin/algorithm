package main

import "fmt"

/*
思路：差分数组+二分查找
时间复杂度：O((m+n)×logn)，其中 n 是 nums 的长度，m 是 queries 的长度。
*/
func minZeroArray(nums []int, queries [][]int) int {
	left, right := 0, len(queries)
	if !check(nums, queries, right) {
		return -1
	}
	for left < right {
		k := (left + right) / 2
		if check(nums, queries, k) {
			right = k
		} else {
			left = k + 1
		}
	}
	return left
}

func check(nums []int, queries [][]int, k int) bool {
	deltaArray := make([]int, len(nums)+1)
	for i := 0; i < k; i++ {
		left, right, value := queries[i][0], queries[i][1], queries[i][2]
		deltaArray[left] += value
		deltaArray[right+1] -= value
	}

	prefixSumArray := make([]int, len(nums)+1)
	var prefixSum int
	for i, delta := range deltaArray {
		prefixSum += delta
		prefixSumArray[i] = prefixSum
	}
	for i := 0; i < len(nums); i++ {
		if prefixSumArray[i] < nums[i] {
			return false
		}
	}

	return true
}

/*
思考：

思路：模拟题
1、
*/
func minZeroArray1(nums []int, queries [][]int) int {
	var zeroCount int
	for _, num := range nums {
		if num == 0 {
			zeroCount++
		}
	}
	if zeroCount == len(nums) {
		return 0
	}

	ans := -1
	for i, query := range queries {
		for j := query[0]; j <= query[1]; j++ {
			if nums[j] == 0 {
				continue
			}
			if nums[j] <= query[2] {
				nums[j] = 0
				zeroCount++
			} else {
				nums[j] -= query[2]
			}
		}
		// fmt.Println("i:", i, ", nums:", nums, ", zeroCount:", zeroCount)

		if zeroCount == len(nums) {
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
