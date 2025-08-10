package main

/*
思路：动态规划
dp[i][j]: 表示nums1[0:i]和nums2[0:j]的最大未交叉线数
状态转移方程：
if nums1[i] == nums2[j]:
	dp[i][j] = dp[i-1][j-1] + 1 (当nums1[i] == nums2[j])
else:
	dp[i][j] = max(dp[i-1][j], dp[i][j-1]) (当nums1[i] != nums2[j])
*/
func maxUncrossedLines1(nums1 []int, nums2 []int) int {
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 == 0 || len2 == 0 {
		return 0
	}

	dp := make([][]int, len1)
	for i := 0; i < len1; i++ {
		dp[i] = make([]int, len2)
	}

	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if nums1[i] == nums2[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1 // 如果是第一个元素，直接赋值为1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1 // 否则加上前一个状态的值
				}
			} else {
				if i != 0 || j != 0 {
					if i == 0 {
						dp[i][j] = dp[i][j-1]
					} else if j == 0 {
						dp[i][j] = dp[i-1][j]
					} else {
						dp[i][j] = max(dp[i-1][j], dp[i][j-1]) // 取前一个状态的最大值
					}
				}
			}
		}
	}

	return dp[len1-1][len2-1]
}

/*
上面的方法对于边界的处理太复杂了，可以简化一下。
*/
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	len1 := len(nums1)
	len2 := len(nums2)

	dp := make([][]int, len1+1)
	for i := 0; i < len1+1; i++ {
		dp[i] = make([]int, len2+1)
	}

	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if nums1[i] == nums2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[len1][len2]
}

func main() {
	testCases := []struct {
		nums1  []int
		nums2  []int
		expect int
	}{
		{
			nums1:  []int{1, 4, 2},
			nums2:  []int{1, 2, 4},
			expect: 2,
		},
	}

	for _, testCase := range testCases {
		result := maxUncrossedLines(testCase.nums1, testCase.nums2)
		if result != testCase.expect {
			println("Test case failed: expected", testCase.expect, "but got", result)
		} else {
			println("Test case passed")
		}
	}
}
