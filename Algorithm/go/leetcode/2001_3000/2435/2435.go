package main

import "fmt"

func numberOfPaths(grid [][]int, k int) int {
	var ans int

	return ans
}

func main() {
	testCases := []struct {
		grid   [][]int
		k      int
		expect int
	}{
		{
			grid:   [][]int{},
			k:      0,
			expect: -2,
		},
	}

	for i, testCase := range testCases {
		ans := numberOfPaths(testCase.grid, testCase.k)
		if ans != testCase.expect {
			fmt.Printf("case %d fail, a:%d, b:%d, expect:%d, ans:%d\n", i, testCase.a, testCase.b, testCase.expect, ans)
		} else {
			fmt.Printf("case %d pass\n", i)
		}
	}
}
