package partition

func lomutoPartition(nums []int, l, r int) int {
	pivot := nums[r] // 选择最右侧元素为 pivot
	i := l - 1       // i 维护小于 pivot 的区域

	for j := l; j < r; j++ { // 遍历 [l, r-1]
		if nums[j] < pivot {
			i++                                 // 扩大小于 pivot 的区域
			nums[i], nums[j] = nums[j], nums[i] // 交换元素
		}
	}

	nums[i+1], nums[r] = nums[r], nums[i+1] // 把 pivot 放到正确位置
	return i + 1                            // 返回 pivot 位置
}
