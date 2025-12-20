package main

func minDeletionSize(strs []string) int {
	if len(strs) == 0 {
		return 0
	}

	strLen := len(strs[0])
	var ans int
	for i := 0; i < strLen; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] < strs[j-1][i] {
				ans++
				break
			}
		}
	}

	return ans
}
