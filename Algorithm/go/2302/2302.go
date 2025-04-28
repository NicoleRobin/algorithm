package main

/*
思路：滑动窗口
1、
*/
func countSubarrays(nums []int, k int64) int64 {
	n := len(nums)
	res, total := int64(0), int64(0)
	for i, j := 0, 0; j < n; j++ {
		total += int64(nums[j])
		for i <= j && total*int64(j-i+1) >= k {
			total -= int64(nums[i])
			i++
		}
		res += int64(j - i + 1)
	}
	return res
}

func main() {
}
