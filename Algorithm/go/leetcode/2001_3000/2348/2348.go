package main

import (
	"fmt"
)

/*
思路：数学
*/
func zeroFilledSubarray1(nums []int) int64 {
	if len(nums) == 0 {
		return 0
	}

	var subArrayLenList []int
	var res int64
	var zeroCount int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
		} else {
			subArrayLenList = append(subArrayLenList, zeroCount)
			zeroCount = 0
		}
	}
	if zeroCount > 0 {
		subArrayLenList = append(subArrayLenList, zeroCount)
	}

	fmt.Println(subArrayLenList)
	for _, count := range subArrayLenList {
		res += int64((1 + count) * count / 2)
	}

	return res
}

/*
思路：
*/
func zeroFilledSubarray(nums []int) int64 {
	var ans int64
	last := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			last = i
		} else {
			ans += int64(i - last)
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums     []int
		expected int64
	}{
		{
			nums:     []int{1, 3, 0, 0, 2, 0, 0, 4},
			expected: 6,
		},
		{
			nums:     []int{0, 0, 0, 2, 0, 0},
			expected: 9,
		},
	}
	for i, tc := range testCases {
		res := zeroFilledSubarray(tc.nums)
		if res != tc.expected {
			fmt.Printf("case #%d fail: got %d, expected %d\n", i+1, res, tc.expected)
		} else {
			fmt.Printf("case #%d success: got %d, expected %d\n", i+1, res, tc.expected)
		}
	}
}
