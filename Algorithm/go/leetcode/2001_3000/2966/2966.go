package main

import (
	"fmt"
	"sort"
)

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)

	var ans [][]int
	for i := 2; i < len(nums); i = i + 3 {
		if nums[i]-nums[i-2] > k {
			ans = [][]int{}
			break
		}
		ans = append(ans, nums[i-2:i+1])
	}
	return ans
}

func main() {
	testCases := []struct {
		nums   []int
		k      int
		expect [][]int
	}{
		{
			nums:   []int{1, 3, 4, 8, 7, 9, 3, 5, 1},
			k:      2,
			expect: [][]int{{1, 1, 3}, {3, 4, 5}, {7, 8, 9}},
		},
	}

	for i, tc := range testCases {
		ans := divideArray(tc.nums, tc.k)
		fmt.Println("i:", i, "ans:", ans)
	}
}
