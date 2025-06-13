package main

/*
思考：
1、先遍历一遍获取最大值；
2、使用滑动窗口算法处理，right指针向右移动，遇到最大值就加1；当统计的最大值个数等于k时，left指针向右移动，直到最大值个数小于k；
3、len(nums) - right和left - lastLeft的乘积就是符合条件的子数组个数；
*/
func countSubarrays(nums []int, k int) int64 {
	if len(nums) == 0 {
		return 0
	}

	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}

	var maxNumCount, left int
	var ans int64
	for right := 0; right < len(nums); right++ {
		if nums[right] == maxNum {
			maxNumCount++
		}
		if maxNumCount == k {
			lastLeft := left
			for left <= right && maxNumCount == k {
				if nums[left] == maxNum {
					maxNumCount--
				}
				left++
			}

			ans += int64(len(nums)-right) * int64(left-lastLeft)
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums   []int
		k      int
		expect int64
	}{
		{
			nums:   []int{1, 3, 2, 3, 3},
			k:      2,
			expect: 6,
		},
		{
			nums:   []int{1, 4, 2, 1},
			k:      3,
			expect: 0,
		},
	}

	for _, testCase := range testCases {
		result := countSubarrays(testCase.nums, testCase.k)
		println("nums:", testCase.nums, "k:", testCase.k, "Result:", result, "Expected:", testCase.expect)
	}
}
