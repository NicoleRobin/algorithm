package main

import "log"

func checkSubarraySum(nums []int, k int) bool {
	preSum := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			preSum[i] = nums[i]
		} else {
			preSum[i] = preSum[i-1] + nums[i]
		}
	}

	preSumResidue := make([]int, len(nums))
	for i := 0; i < len(preSum); i++ {
		preSumResidue[i] = preSum[i] % k
	}
	log.Println(preSumResidue)

	mp := map[int][]int{
		0: {-1},
	}
	exist := false
	for i := 0; i < len(preSumResidue); i++ {
		if indexList, ok := mp[preSumResidue[i]]; ok {
			for _, index := range indexList {
				if i-index >= 2 {
					exist = true
				}
			}
			mp[preSumResidue[i]] = append(indexList, i)
		} else {
			mp[preSumResidue[i]] = []int{i}
		}

		if exist {
			break
		}
	}
	return exist
}

func checkSubarraySum2(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}

	mp := map[int]int{
		0: -1,
	}
	remainder := 0
	exist := false
	for i, num := range nums {
		remainder = (remainder + num) % k
		if previousIndex, ok := mp[remainder]; ok {
			if i-previousIndex >= 2 {
				exist = true
			}
		} else {
			mp[remainder] = i
		}
	}
	return exist
}

func main() {
	nums := []int{23, 2, 4, 6, 6}
	k := 7

	result := checkSubarraySum(nums, k)
	log.Println(result)
}
