package main

import (
	"fmt"
	"math"
)

func divide(a int, b int) int {
	// 特殊情况处理
	if b == 0 {
		panic("division by zero")
	}
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}

	// 确定结果符号
	negative := (a < 0) != (b < 0)

	// 转为负数，防止溢出
	aa, bb := -abs(a), -abs(b)

	result := 0
	for aa <= bb {
		tmp, multiple := bb, 1
		// 尝试将 bb 成倍扩大
		for aa <= (tmp<<1) && (tmp<<1) < 0 {
			tmp <<= 1
			multiple <<= 1
		}
		aa -= tmp
		result += multiple
	}

	if negative {
		return -result
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	testCases := []struct {
		a, b, expect int
	}{
		{
			a:      7,
			b:      -3,
			expect: -2,
		},
	}

	for i, testCase := range testCases {
		ans := divide(testCase.a, testCase.b)
		if ans != testCase.expect {
			fmt.Printf("case %d fail, a:%d, b:%d, expect:%d, ans:%d\n", i, testCase.a, testCase.b, testCase.expect, ans)
		} else {
			fmt.Printf("case %d pass\n", i)
		}
	}
}
