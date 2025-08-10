package main

import "fmt"

func countGood(nums []int, k int) int64 {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, 0
	numCountMap := map[int]int{
		nums[0]: 1,
	}

	var result int64
	for right < len(nums) {
		var goodArrayCount int
		for _, count := range numCountMap {
			goodArrayCount += count * (count - 1) / 2
		}

		fmt.Println("left:", left, ", right:", right, ", goodArrayCount:", goodArrayCount, "subArray:", nums[left:right+1], ", numCountMap:", numCountMap)
		if goodArrayCount >= k {
			// 如果当前满足要求，那么加上所有的右边元素的子数组都满足要求
			result += int64(len(nums) - right)

			numCountMap[nums[left]]--
			left++
		} else {
			right++
			if right >= len(nums) {
				break
			}
			numCountMap[nums[right]]++
		}
	}

	return result
}

func countGood1(nums []int, k int) int64 {
	n := len(nums)
	same, right := 0, -1
	cnt := make(map[int]int)
	var ans int64 = 0
	for left := 0; left < n; left++ {
		for same < k && right+1 < n {
			right++
			same += cnt[nums[right]]
			cnt[nums[right]]++
		}
		if same >= k {
			ans += int64(n - right)
		}
		cnt[nums[left]]--
		same -= cnt[nums[left]]
	}
	return ans
}

func main() {
	nums := []int{3, 1, 4, 3, 2, 2, 4}
	result := countGood(nums, 2)
	fmt.Println(result)
}
