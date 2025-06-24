package main

func findKDistantIndices(nums []int, key int, k int) []int {
	var ans []int
	var validStart, validEnd int
	for i := 0; i < len(nums); i++ {
		if nums[i] == key {
			if validStart < i-k {
				validStart = i - k
			}

			validEnd = len(nums)
			if validEnd > i+k+1 {
				validEnd = i + k + 1
			}

			for j := validStart; j < validEnd; j++ {
				ans = append(ans, j)
			}

			validStart = validEnd
		}
	}

	return ans
}

func main() {

}
