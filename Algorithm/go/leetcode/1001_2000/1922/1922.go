package main

import (
	"fmt"
)

/*
质数：2 3 5 7
偶数：0 2 4 6 8

复盘：
1、该题其实是要考察快速幂算法
*/
func countGoodNumbers(n int64) int {
	var mod int64 = 1000000007

	// 快速幂算法
	quickmul := func(x, y int64) int64 {
		ret := int64(1)
		mul := x
		for y > 0 {
			if y%2 == 1 {
				ret = ret * mul % mod
			}
			mul = mul * mul % mod
			y /= 2
		}
		return ret
	}

	return int(quickmul(5, (n+1)/2) * quickmul(4, n/2) % mod)
}

func main() {
	result := countGoodNumbers(60)
	fmt.Println(result)
}
