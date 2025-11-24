package main

import "fmt"

/*
思路：
1、没有很好的思路，先实现暴力算法
*/
func maxSumDivThree1(nums []int) int {
	var ans int

	return ans
}

/*
官方题解：动态规划，双层转移方程
1、状态定义：
f(i,j): 在前i个数中选取了若干个数，并且他们的和模3为j时，这些数的和的最大值；
2、状态转移方程：
f(i,j)=max{f(i−1,j),f(i−1,(j−nums[i])mod3)+nums[i]}
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSumDivThree(nums []int) int {
	f := []int{0, -0x3f3f3f3f, -0x3f3f3f3f}
	for _, num := range nums {
		g := make([]int, 3)
		for i := 0; i < 3; i++ {
			g[(i+num)%3] = max(f[(i+num)%3], f[i]+num)
		}
		f = g
	}
	return f[0]
}

func main() {
	testCases := []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{1, 2, 3},
			expect: -2,
		},
	}

	for i, testCase := range testCases {
		ans := maxSumDivThree(testCase.nums)
		if ans != testCase.expect {
			fmt.Printf("case %d fail, nums:%d, expect:%d, ans:%d\n", i, testCase.nums, testCase.expect, ans)
		} else {
			fmt.Printf("case %d pass\n", i)
		}
	}
}
