package main

/*
思考：
1、要获取最大值，只能往前跳，因为往后跳的条件是后面的值小于当前值；
2、还要考虑可以先往后跳，再往前跳的情况，可以获得更大的值；
3、但是要寻找最大值最终还是必须要往前跳，往后跳只是可以把位置往后挪动一些，越往后的位置越可能找到更大的值；
*/
func maxValue(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	ans := make([]int, len(nums))
	ans[0] = nums[0]
	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
			ans[i] = maxNum
		} else {
			ans[i] = maxNum
		}
	}

	// 再考虑往后跳的情况
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				ans[i] = max(ans[i], ans[j])
			}
		}
	}

	return ans
}

func main() {}
