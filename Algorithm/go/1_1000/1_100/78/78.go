package main

import "fmt"

func subsets(nums []int) [][]int {
	var result [][]int
	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)

		if start == len(nums) {
			return
		}

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i+1, path)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, []int{})

	return result
}

type TestCase struct {
	nums []int
}

func main() {
	fmt.Println("leetcode 78")
	TestCases := []TestCase{
		{[]int{1, 2, 3}},
	}
	for _, testCase := range TestCases {
		result := subsets(testCase.nums)
		for _, row := range result {
			fmt.Println("row:", row)
		}
	}
}
