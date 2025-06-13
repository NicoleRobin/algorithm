package main

/*
思路：位运算

*/
func singleNumber(nums []int) []int {
	xorResult := 0
	for _, num := range nums {
		xorResult ^= num // 异或操作，两个相同的数异或结果为0
	}

	// 找到异或结果的最低位1
	lsb := xorResult & -xorResult
	a, b := 0, 0
	for _, num := range nums {
		if num&lsb > 0 {
			a ^= num // 如果最低位1在num上，则将num加入a
		} else {
			b ^= num // 否则将num加入b
		}
	}

	return []int{a, b}
}

func main() {
}
