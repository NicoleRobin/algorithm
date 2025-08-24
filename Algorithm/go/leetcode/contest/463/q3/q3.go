package main

/*
思路：贪心？

思考：
1、不同的处理顺序会不会影响最终结果呢？

重点：
1、连续的子数组
*/
func minArraySum(nums []int, k int) int64 {
	var res int64
	remain := len(nums)
	for remain >= k {
		var presum int
		presums := make([]int, len(nums)+1)
		for i := 0; i < len(nums); i++ {
			presum += nums[i]
			presums[i+1] = presum
		}

		for i := 0; i+k < len(nums); i++ {
			if presums[i+k]-presums[i]%k == 0 {
				for j := i; j < i+k; j++ {
					remain--
					nums[j] = 0
				}
			}
		}
	}

	return res
}

func main() {}
