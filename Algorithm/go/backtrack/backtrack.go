package backtrack

func backtrack(nums []int, start int, path []int, res *[][]int) {
	if start == len(nums) {
		*res = append(*res, append([]int{}, path...))
		return
	}

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		backtrack(nums, i+1, path, res)
		path = path[:len(path)-1]
	}
}
