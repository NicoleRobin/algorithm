package main

import (
	"math"
	"sort"
)

/*
思路：
1、贪心

将nums排序，然后依次选取可选的最小值
*/
func maxDistinctElements(nums []int, k int) int {
	var ans int
	sort.Ints(nums)
	prev := math.MinInt32
	for _, num := range nums {
		curr := min(max(num-k, prev+1), num+k)
		if curr > prev {
			ans++
			prev = curr
		}
	}

	return ans
}

func main() {

}
