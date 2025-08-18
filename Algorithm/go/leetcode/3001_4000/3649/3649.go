package main

import "slices"

/*
思路：暴力法
*/
func perfectPairs(nums []int) int64 {
	var res int64
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			aMinusb := abs(nums[i] - nums[j])
			aPlusb := abs(nums[j] + nums[i])
			minAb := min(abs(nums[i]), abs(nums[j]))
			maxAb := max(abs(nums[j]), abs(nums[i]))
			if min(aMinusb, aPlusb) <= minAb && max(aMinusb, aPlusb) >= maxAb {
				res++
			}
		}
	}

	return res
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

/*
思路：
*/
func perfectPairs1(nums []int) (ans int64) {
	for i, x := range nums {
		if x < 0 {
			nums[i] *= -1
		}
	}

	slices.Sort(nums)
	left := 0
	for j, b := range nums {
		for nums[left]*2 < b {
			left++
		}
		// a=nums[i]，其中 i 最小是 left，最大是 j-1，一共有 j-left 个
		ans += int64(j - left)
	}
	return
}

func main() {}
