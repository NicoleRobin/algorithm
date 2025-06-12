package main

func maxAdjacentDistance(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var ans int
	var diff int
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			diff = nums[i] - nums[i-1]
		} else {
			diff = nums[i-1] - nums[i]
		}
		if diff > ans {
			ans = diff
		}
	}
	if nums[len(nums)-1] > nums[0] {
		diff = nums[len(nums)-1] - nums[0]
	} else {
		diff = nums[0] - nums[len(nums)-1]
	}
	if diff > ans {
		ans = diff
	}

	return ans
}

func main() {

}
