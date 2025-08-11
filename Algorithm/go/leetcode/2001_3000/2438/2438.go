package main

import (
	"fmt"
	"reflect"
)

/*
思路：模拟题
1、先将n拆解为不同的2的幂
2、然后按照要求求出对应的answer
*/
func productQueries(n int, queries [][]int) []int {
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
			tmp %= 1e9
		}
		res = append(res, tmp)
	}
	return res
}

func main() {
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
	}
	for i, tc := range testCases {
		res := productQueries(tc.n, tc.queries)
		if reflect.DeepEqual(res, tc.expected) {
			fmt.Printf("Case %d fail, expected %d got %d\n", i, tc.expected, res)
		} else {
			fmt.Printf("Case %d pass, got %d\n", i, res)
		}
	}
}
