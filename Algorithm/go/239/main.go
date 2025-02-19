package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	result := maxSlidingWindow(nums, 3)
	fmt.Println("result:", result)
}

type numIndex struct {
	Num   int
	Index int
}

func maxSlidingWindow(nums []int, k int) []int {
	var queue []numIndex
	var result []int
	for i := 0; i < len(nums); i++ {
		// fmt.Println("i:", i, ", queue:", queue)
		for len(queue) > 0 {
			if queue[len(queue)-1].Num < nums[i] {
				queue = queue[:len(queue)-1]
			} else {
				break
			}
		}
		queue = append(queue, numIndex{Num: nums[i], Index: i})

		// 去除队头已经失效的值
		var validStart int
		for validStart = 0; validStart < len(queue); validStart++ {
			if queue[validStart].Index > i-k {
				break
			}
		}
		queue = queue[validStart:]

		if i >= k-1 && len(queue) > 0 {
			result = append(result, queue[0].Num)
		}
	}
	return result
}
