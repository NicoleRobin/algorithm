package main

import "sort"

func partitionArray(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	var ans int
	start := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[start] > k {
			ans++
			start = i
		}
	}
	if start < len(nums) {
		ans++
	}

	return ans
}

func main() {

}
