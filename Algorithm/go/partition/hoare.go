package partition

func hoarePartition(nums []int, l, r int) int {
	pivot := nums[l] // 选择左端点为 pivot
	i, j := l-1, r+1 // 关键初始化

	for {
		// 找到左侧第一个 >= pivot 的元素
		for i++; nums[i] < pivot; i++ {
		}

		// 找到右侧第一个 <= pivot 的元素
		for j--; nums[j] > pivot; j-- {
		}

		if i >= j {
			return j // j 是分区边界
		}

		// 交换 nums[i] 和 nums[j]
		nums[i], nums[j] = nums[j], nums[i]
	}
}
