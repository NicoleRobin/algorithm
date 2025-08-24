package main

import "fmt"

/*
思路：
1、尝试删除每个字符然后计算全为1的子数组长度，求最大值；
时间复杂度：O(n*n)

优化思路：滑动窗口，当窗口中的0个数大于1时移动左侧，否则移动右侧
*/
func longestSubarray1(nums []int) int {
	var ans int
	for i := 0; i < len(nums); i++ {
		var oneArrayLen int
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			if nums[j] == 1 {
				oneArrayLen++
				if oneArrayLen > ans {
					ans = oneArrayLen
				}
			} else {
				oneArrayLen = 0
			}
		}
	}

	return ans
}

/*
思考：
1、虽然最终通过各种优化解决了这个问题，但是做题的时候还是总是在尝试，无法在最一开始想到滑动窗口方案时就把思路想明白，还需要继续学习呀；
2、
*/
func longestSubarray(nums []int) int {
	var ans int
	var left int
	var zeroCount int
	var firstZeroIndex int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
			if zeroCount == 1 {
				firstZeroIndex = i
			}
			if zeroCount > 1 {
				left = firstZeroIndex + 1
				zeroCount--
				firstZeroIndex = i
			}
		}
		if i-left > ans {
			ans = i - left
		}
		fmt.Printf("i: %d, left: %d, zeroCount: %d, ans:%d\n", i, left, zeroCount, ans)
	}

	return ans
}

func main() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		/*
			{
				nums:     []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
				expected: 5,
			},
			{
				nums:     []int{1, 1, 0, 1},
				expected: 3,
			},

		*/
		{
			nums:     []int{1, 1, 0, 0, 1, 1, 1, 0, 1},
			expected: 4,
		},
	}
	for i, testCase := range testCases {
		result := longestSubarray(testCase.nums)
		if result != testCase.expected {
			fmt.Printf("Case %d fail, expected:%d, ans:%d\n", i, testCase.expected, result)
		} else {
			fmt.Printf("Case %d pass\n", i)
		}
	}
}
