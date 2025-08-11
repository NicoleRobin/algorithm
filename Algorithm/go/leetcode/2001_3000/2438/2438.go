package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
思路：模拟题
1、先将n拆解为不同的2的幂
2、然后按照要求求出对应的answer

优化思路：
1、前缀积
*/
func productQueries(n int, queries [][]int) []int {
	const mod = 1000000007
	var bins []int
	rep := 1
	for n > 0 {
		if n%2 == 1 {
			bins = append(bins, rep)
		}
		n /= 2
		rep *= 2
	}

	var res []int
	for _, query := range queries {
		tmp := 1
		for i := query[0]; i <= query[1]; i++ {
			tmp *= bins[i]
			tmp %= mod
		}
		res = append(res, tmp)
	}
	return res
}

func calcBins(n int) []int {
	var bins []int
	rep := 1
	for n > 0 {
		if n%2 == 1 {
			bins = append(bins, rep)
		}
		n /= 2
		rep *= 2
	}
	return bins
}

func calcBins1(n int) []int {
	binN := fmt.Sprintf("%b", n)

	var res []int
	for i := 0; i < len(binN); i++ {
		if binN[len(binN)-1-i] == '1' {
			res = append(res, int(math.Pow(2, float64(i))))
		}
	}
	return res
}

func main() {
	bins := calcBins(919)
	fmt.Printf("bins:%+v, mod:%d\n", bins, 1e9+7)

	testCases := []struct {
		n        int
		queries  [][]int
		expected []int
	}{
		{
			n:        15,
			queries:  [][]int{{0, 1}, {2, 2}, {0, 3}},
			expected: []int{2, 4, 64},
		},
		{
			n:        919,
			queries:  [][]int{{5, 5}, {4, 4}, {0, 1}, {1, 5}, {4, 6}, {6, 6}, {5, 6}, {0, 3}, {5, 5}, {5, 6}, {1, 2}, {3, 5}, {3, 6}, {5, 5}, {4, 4}, {1, 1}, {2, 4}, {4, 5}, {4, 4}, {5, 6}, {0, 4}, {3, 3}, {0, 4}, {0, 5}, {4, 4}, {5, 5}, {4, 6}, {4, 5}, {0, 4}, {6, 6}, {6, 6}, {6, 6}, {2, 2}, {0, 5}, {1, 4}, {0, 3}, {2, 4}, {5, 5}, {6, 6}, {2, 2}, {2, 3}, {5, 5}, {0, 6}, {3, 3}, {6, 6}, {4, 4}, {0, 0}, {0, 2}, {6, 6}, {6, 6}, {3, 6}, {0, 4}, {6, 6}, {2, 2}, {4, 6}},
			expected: []int{256, 128, 2, 4194304, 16777216, 512, 131072, 128, 256, 131072, 8, 524288, 268435456, 256, 128, 2, 8192, 32768, 128, 131072, 16384, 16, 16384, 4194304, 128, 256, 16777216, 32768, 16384, 512, 512, 512, 4, 4194304, 16384, 128, 8192, 256, 512, 4, 64, 256, 147483634, 16, 512, 128, 1, 8, 512, 512, 268435456, 16384, 512, 4, 16777216},
		},
	}
	for i, tc := range testCases {
		res := productQueries(tc.n, tc.queries)
		if !reflect.DeepEqual(res, tc.expected) {
			fmt.Printf("Case %d fail, expected %d got %d\n", i, tc.expected, res)
		} else {
			fmt.Printf("Case %d pass, got %d\n", i, res)
		}
	}
}
