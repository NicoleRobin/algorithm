package main

import (
	"fmt"
	"math"
)

/*
x：被除数
y：除数
z：
这里是要判断divisor*mid是大于还是小于dividend
*/
func quickAdd(divisor, mid, dividend int) bool {
	fmt.Printf("divisor:%d, mid:%d, dividend:%d\n", divisor, mid, dividend)
	for result, add := 0, divisor; mid > 0; mid >>= 1 { // 不能使用除法
		fmt.Printf("result:%d, add:%d, mid:%b\n", result, add, mid)
		if mid&1 > 0 {
			// 需要保证 result + add >= x
			if result < dividend-add {
				return false
			}
			result += add
		}
		if mid != 1 {
			// 需要保证 add + add >= x
			if add < dividend-add {
				return false
			}
			add += add
		}
	}
	return true
}

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 { // 考虑被除数为最小值的情况
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 { // 考虑除数为最小值的情况
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend == 0 { // 考虑被除数为 0 的情况
		return 0
	}

	// 一般情况，使用二分查找
	// 将所有的正数取相反数，这样就只需要考虑一种情况
	rev := false
	if dividend > 0 {
		dividend = -dividend
		rev = !rev
	}
	if divisor > 0 {
		divisor = -divisor
		rev = !rev
	}

	ans := 0
	left, right := 1, math.MaxInt32
	for left <= right {
		mid := left + (right-left)>>1 // 注意溢出，并且不能使用除法
		if quickAdd(divisor, mid, dividend) {
			ans = mid
			if mid == math.MaxInt32 { // 注意溢出
				break
			}
			left = mid + 1
			fmt.Println("in right")
		} else {
			right = mid - 1
			fmt.Println("in left")
		}
	}
	if rev {
		return -ans
	}
	return ans
}

func main() {
	fmt.Println("leetcode")
	testCases := [][]int{
		{-7, 3},
		// {10, 3},
	}
	for _, testCase := range testCases {
		result := divide(testCase[0], testCase[1])
		fmt.Printf("dividend:%d, divisor:%d, result:%d\n", testCase[0], testCase[1], result)
	}

	// fmt.Printf("-1 << 31:%d, 1 >> 31 - 1:%d\n", -1<<31, 1<<31-1)
}
