package main

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	len1 := len(nums1)
	len2 := len(nums2)
	dp := make([][]int, len1)
	for i := 0; i < len1; i++ {
		dp[i] = make([]int, len2)
	}

	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if nums1[i] == nums2[j] {

			} else {

			}
		}
	}

	return dp[len1-1][len2-1]
}

func main() {

}
