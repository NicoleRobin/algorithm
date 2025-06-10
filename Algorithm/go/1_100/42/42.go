package main

import "fmt"

/*
思路：
1、遍历数组
2、对于每个位置计算其可以接到的雨水数，计算规则：max(左侧最大值, 右侧最大值) - 当前高度值

时间复杂度：O(n^2)
外层循环遍历一遍数组，对于每一个位置，最差需要遍历整个数组找到左右最大值
*/
func trap1(height []int) int {
	result := 0
	for i := 1; i < len(height)-1; i++ {
		left, right := 0, 0
		for j := i - 1; j >= 0; j-- {
			if height[j] > left {
				left = height[j]
			}
		}
		for j := i + 1; j < len(height); j++ {
			if height[j] > right {
				right = height[j]
			}
		}

		max := 0
		if left > right {
			max = right
		} else {
			max = left
		}

		if max > height[i] {
			result += max - height[i]
		}
	}
	return result
}

/*
思路：
1、先通过动态规划思路计算每一个位置左侧最大值以及右侧最大值
2、遍历一次数组

时间复杂度：O(n)
通过空间换时间的思路将时间复杂度提升到O(n)
*/
func trap2(height []int) int {
	if len(height) == 0 {
		return 0
	}

	dpLeft := make([]int, len(height))
	dpRight := make([]int, len(height))
	leftMax := height[0]
	dpLeft[0] = leftMax
	for i := 1; i < len(height); i++ {
		if height[i-1] > leftMax {
			leftMax = height[i-1]
		}
		dpLeft[i] = leftMax
	}
	fmt.Println(dpLeft)

	rightMax := height[len(height)-1]
	dpRight[len(height)-1] = rightMax
	for i := len(height) - 2; i >= 0; i-- {
		if height[i+1] > rightMax {
			rightMax = height[i+1]
		}
		dpRight[i] = rightMax
	}
	fmt.Println(dpRight)

	result := 0
	for i := 1; i < len(height)-1; i++ {
		cur := 0
		if dpLeft[i] > dpRight[i] {
			cur = dpRight[i] - height[i]
		} else {
			cur = dpLeft[i] - height[i]
		}
		if cur > 0 {
			result += cur
		}
	}
	return result
}

/*
思路：
1、通过双指针优化trap2方案的空间复杂度
2、某个位置能接到的雨水数取决于左右两边最大值中的较小者，因此每次不需要都计算出两边的最大值，
只要某一边最大值小于另一边的任意值，那么最终较小者一定是第一个最大值，因此可以只用两个变量保存两边的最大值，交替计算。

时间复杂度：O(n)
空间复杂度：O(1)
*/
func trap3(height []int) int {
	if len(height) == 0 {
		return 0
	}

	result := 0
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left <= right {
		// 这里是应该<呢还是<=呢？
		// 最终应该是根据left==right时是否有意义，有意义就<=，否则就<
		if leftMax <= rightMax {
			if leftMax > height[left] {
				result += leftMax - height[left]
			}
			if height[left] > leftMax {
				leftMax = height[left]
			}
			left++
		} else {
			if rightMax > height[right] {
				result += rightMax - height[right]
			}
			if height[right] > rightMax {
				rightMax = height[right]
			}
			right--
		}
	}
	return result
}

func main() {
	fmt.Println("vim-go")
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	result := trap1(height)
	fmt.Println(result)
}
