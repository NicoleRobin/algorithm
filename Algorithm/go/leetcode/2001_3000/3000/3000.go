package main

import "fmt"

/*
反思：
1、本来是很简单的题，但是因为没有注意到其中一句话：如果存在多个对角线长度相同的矩形，返回面积最 大 的矩形的面积。
导致忽略了很多细节，所以错了很多次，以后要注意啦。
*/
func areaOfMaxDiagonal(dimensions [][]int) int {
	var maxDimension int
	var ans int
	for i := 0; i < len(dimensions); i++ {
		dimension := dimensions[i][0]*dimensions[i][0] + dimensions[i][1]*dimensions[i][1]
		fmt.Printf("i:%d, dimension:%d, dimensions:%v\n", i, dimension, dimensions[i])
		if dimension > maxDimension {
			maxDimension = dimension
			ans = dimensions[i][0] * dimensions[i][1]
		} else if dimension == maxDimension {
			if dimensions[i][0]*dimensions[i][1] > ans {
				ans = dimensions[i][0] * dimensions[i][1]
			}
		}
	}
	return ans
}

func main() {
	testCases := []struct {
		dimensions [][]int
		expected   int
	}{
		{
			dimensions: [][]int{{25, 60}, {39, 52}, {16, 63}, {33, 56}},
			expected:   2028,
		},
		{
			dimensions: [][]int{{6, 5}, {8, 6}, {2, 10}, {8, 1}, {9, 2}, {3, 5}, {3, 5}},
			expected:   20,
		},
	}
	for i, tc := range testCases {
		ans := areaOfMaxDiagonal(tc.dimensions)
		if ans != tc.expected {
			fmt.Printf("Case: %d fail, expected:%d, ans:%d\n", i, tc.expected, ans)
		} else {
			fmt.Printf("Case: %d pass\n", i)
		}
	}
}
